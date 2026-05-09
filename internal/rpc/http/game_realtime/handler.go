package game_realtime

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	ws "github.com/gorilla/websocket"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/uuidcheck"
)

type actionRunner interface {
	Run(ctx context.Context, r *models.Room, action *models.GameAction) (*models.GameUpdate, error)
}

type Handler struct {
	hub      *Hub
	runner   actionRunner
	upgrader ws.Upgrader
	logger   *zap.Logger
}

func NewHandler(hub *Hub, runner actionRunner, allowedOrigins []string, logger *zap.Logger) *Handler {
	return &Handler{
		hub:    hub,
		runner: runner,
		logger: logger,
		upgrader: ws.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				origin := r.Header.Get("Origin")
				for _, allowed := range allowedOrigins {
					if strings.EqualFold(strings.TrimSpace(allowed), origin) {
						return true
					}
				}
				return false
			},
		},
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user")
	gameID := r.URL.Query().Get("game")
	if err := uuidcheck.Check(userID, gameID); err != nil {
		h.logger.Error("invalid uuid", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.logger.Error("upgrade failed", zap.Error(err))
		return
	}
	defer conn.Close()

	player := &models.Player{Id: userID, Conn: conn}
	room, err := h.hub.GetOrCreate(r.Context(), gameID)
	if err != nil {
		h.logger.Error("get/create room failed", zap.Error(err))
		return
	}

	switch {
	case room.Player1 == nil || room.Player1.Id == player.Id:
		room.Player1 = player
	case room.Player2 == nil || room.Player2.Id == player.Id:
		room.Player2 = player
	default:
		h.logger.Error("third connection rejected",
			zap.Any("room", room), zap.String("user", player.Id))
		return
	}

	defer h.hub.Leave(gameID, userID)

	go h.pinger(r.Context(), conn)

	if err := conn.SetReadDeadline(time.Now().Add(models.PongWait)); err != nil {
		h.logger.Error("set read deadline failed", zap.Error(err))
		return
	}
	conn.SetPongHandler(player.PongHandler)

	for {
		var action models.GameAction
		_, msg, err := conn.ReadMessage()
		if err != nil {
			if ws.IsCloseError(err, ws.CloseNoStatusReceived, ws.CloseNormalClosure, ws.CloseAbnormalClosure) {
				return
			}
			h.logger.Error("read message failed", zap.Error(err))
			continue
		}
		if err := json.Unmarshal(msg, &action); err != nil {
			h.logger.Error("decode action failed", zap.Error(err))
			continue
		}
		if uuidcheck.Check(action.UserID) != nil || len(action.Numbers) != 16 {
			h.logger.Error("invalid action payload",
				zap.Any("action", action))
			continue
		}

		update, err := h.runner.Run(r.Context(), room, &action)
		if err != nil {
			h.logger.Error("run action failed", zap.Error(err))
		}
		if update == nil {
			continue
		}

		switch update.Status {
		case models.GameStart, models.GameEnd:
			h.notifyAll(room, update)
			if update.Status == models.GameEnd {
				if room.Player1 != nil && room.Player1.Conn != nil {
					_ = room.Player1.Conn.Close()
				}
				if room.Player2 != nil && room.Player2.Conn != nil {
					_ = room.Player2.Conn.Close()
				}
				return
			}
		default:
			h.notifyOpponent(room, player, update)
		}
	}
}

func (h *Handler) notifyAll(r *models.Room, update *models.GameUpdate) {
	for _, p := range []*models.Player{r.Player1, r.Player2} {
		if p == nil {
			continue
		}
		if err := p.Conn.WriteJSON(update); err != nil {
			h.logger.Error("notify all", zap.Error(err))
		}
	}
}

func (h *Handler) notifyOpponent(r *models.Room, sender *models.Player, update *models.GameUpdate) {
	opp := r.Player1
	if sender == opp {
		opp = r.Player2
	}
	if opp == nil {
		return
	}
	if err := opp.Conn.WriteJSON(update); err != nil {
		h.logger.Error("notify opponent", zap.Error(err))
	}
}

func (h *Handler) pinger(ctx context.Context, conn *ws.Conn) {
	t := time.NewTicker(models.PingInterval)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			if err := conn.WriteMessage(ws.PingMessage, nil); err != nil {
				h.logger.Error("ping failed", zap.Error(err))
				return
			}
		}
	}
}
