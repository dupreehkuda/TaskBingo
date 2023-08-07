package handlers

import (
	"context"
	"net/http"
	"time"

	ws "github.com/gorilla/websocket"
	"github.com/mailru/easyjson"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

var wsUpgrade = ws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// getOrCreateRoom retrieves game from the hub or repository
func (h *handlers) getOrCreateRoom(ctx context.Context, gameID string) (*models.Room, error) {
	h.hub.Mu.Lock()
	defer h.hub.Mu.Unlock()

	if game, ok := h.hub.Rooms[gameID]; ok {
		return game, nil
	}

	game, err := h.service.GetRoom(ctx, gameID)
	if err != nil {
		h.logger.Error("Error occurred fetching game", zap.Error(err))
		return nil, err
	}

	h.hub.Rooms[gameID] = game

	return game, nil
}

// removeRoom removes room from the hub
func (h *handlers) removeRoom(gameID string) {
	h.hub.Mu.Lock()
	defer h.hub.Mu.Unlock()

	if _, ok := h.hub.Rooms[gameID]; !ok {
		return
	}

	delete(h.hub.Rooms, gameID)
}

func (h *handlers) leaveRoom(gameID, userID string) {
	h.hub.Mu.Lock()
	defer h.hub.Mu.Unlock()

	if h.hub.Rooms[gameID] != nil {
		if h.hub.Rooms[gameID].Player1 != nil {
			if h.hub.Rooms[gameID].Player1.Id == userID {
				h.hub.Rooms[gameID].Player1 = nil
			}
		}

		if h.hub.Rooms[gameID].Player2 != nil {
			if h.hub.Rooms[gameID].Player2.Id == userID {
				h.hub.Rooms[gameID].Player2 = nil
			}
		}
	}

	return
}

// notifyOpponent notifies opponent with event
func (h *handlers) notifyOpponent(room *models.Room, sender *models.Player, update *models.GameUpdate) error {
	opponent := room.Player1
	if sender == opponent {
		opponent = room.Player2
	}

	if opponent == nil {
		h.logger.Debug("there is no opponent")
		return nil
	}

	err := opponent.Conn.WriteJSON(update)
	if err != nil {
		h.logger.Error("Error notifying player", zap.Error(err))
		return err
	}

	return nil
}

// notifyAll notifies all players in a room with event
func (h *handlers) notifyAll(room *models.Room, update *models.GameUpdate) error {
	if room.Player1 != nil {
		if err := room.Player1.Conn.WriteJSON(update); err != nil {
			h.logger.Error("Error notifying player", zap.Error(err))
			return err
		}
	}

	if room.Player2 != nil {
		if err := room.Player2.Conn.WriteJSON(update); err != nil {
			h.logger.Error("Error notifying player", zap.Error(err))
			return err
		}
	}

	return nil
}

// GameWSLaunch upgrades ws connection to handle active game requests
func (h *handlers) GameWSLaunch(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user")
	gameID := r.URL.Query().Get("game")

	if err := UUIDCheck(userID, gameID); err != nil {
		h.logger.Error("Invalid UUID in request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	conn, err := wsUpgrade.Upgrade(w, r, nil)
	if err != nil {
		h.logger.Error("Error upgrading connection", zap.Error(err))
		return
	}
	defer conn.Close()

	player := &models.Player{
		Id:   userID,
		Conn: conn,
	}

	room, err := h.getOrCreateRoom(r.Context(), gameID)
	if err != nil {
		h.logger.Error("Error getting game", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch {
	case room.Player1 == nil || room.Player1.Id == player.Id:
		room.Player1 = player
	case room.Player2 == nil || room.Player2.Id == player.Id:
		room.Player2 = player
	default:
		h.logger.Error("Error third connection", zap.Any("gameInfo", room), zap.String("thirdUser", player.Id))
		return
	}

	defer func() {
		h.leaveRoom(gameID, userID)

		if room.Player1 == nil && room.Player2 == nil {
			h.removeRoom(room.Id)
		}
	}()

	go func() {
		err := h.Pinger(conn, r.Context())
		if err != nil {
			h.logger.Error("pinger error", zap.Error(err))
			return
		}
	}()

	if err := conn.SetReadDeadline(time.Now().Add(models.PongWait)); err != nil {
		h.logger.Error("Unable to set read deadline", zap.Error(err))
		return
	}

	conn.SetPongHandler(player.PongHandler)

	for {
		var action models.GameAction
		_, msg, err := conn.ReadMessage()
		if err != nil {
			if ws.IsCloseError(err, ws.CloseNoStatusReceived, ws.CloseNormalClosure) {
				break
			}

			h.logger.Error("Unable to read message", zap.Error(err))
			continue
		}

		if err = easyjson.Unmarshal(msg, &action); err != nil {
			h.logger.Error("Unable to decode JSON", zap.Error(err))
			continue
		}

		if err = UUIDCheck(action.UserID); err != nil || len(action.Numbers) != 16 {
			h.logger.Error("Invalid UUID or missing numbers", zap.Error(err))
			continue
		}

		update, err := h.service.UpdateGame(r.Context(), room, &action)
		if err != nil {
			h.logger.Error("Error getting update", zap.Error(err))
		}

		if update == nil {
			continue
		}

		switch update.Status {
		case models.GameStart:
			if err = h.notifyAll(room, update); err != nil {
				h.logger.Error("Error on notifying player", zap.Error(err))
			}
			continue

		case models.GameEnd:
			if err = h.notifyAll(room, update); err != nil {
				h.logger.Error("Error on notifying player", zap.Error(err))
			}
			return

		default:
			if err = h.notifyOpponent(room, player, update); err != nil {
				h.logger.Error("Error on notifying player", zap.Error(err))
			}
		}
	}

	h.logger.Debug("Connection closed", zap.String("userID", userID))
}

func (h *handlers) Pinger(conn *ws.Conn, ctx context.Context) error {
	ticker := time.NewTicker(models.PingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			h.logger.Debug("ctx done")
			return nil

		case <-ticker.C:
			if err := conn.WriteMessage(ws.PingMessage, []byte{}); err != nil {
				h.logger.Error("unable to ping client", zap.Error(err))
				return err
			}
		}
	}
}

// pongHandler is used to handle PongMessages for the Client
func (h *handlers) pongHandler(conn *ws.Conn) error {
	return conn.SetReadDeadline(time.Now().Add(models.PongWait))
}
