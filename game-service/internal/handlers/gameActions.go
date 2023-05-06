package handlers

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

type (
	Player struct {
		Id   string `json:"id"`
		Conn *websocket.Conn
	}

	Game struct {
		Id      string  `json:"id"`
		Player1 *Player `json:"player1"`
		Player2 *Player `json:"player2"`
	}

	GameHub struct {
		mu    sync.Mutex
		games map[string]*Game
	}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *handlers) getOrCreateGame(id string) *Game {
	h.hub.mu.Lock()
	defer h.hub.mu.Unlock()

	if game, ok := h.hub.games[id]; ok {
		return game
	}

	game := &Game{Id: id}
	h.hub.games[id] = game

	return game
}

func (h *handlers) removeRoom(id string) {
	h.hub.mu.Lock()
	defer h.hub.mu.Unlock()

	if _, ok := h.hub.games[id]; !ok {
		return
	}

	delete(h.hub.games, id)
}

func (h *handlers) notifyOpponent(game *Game, sender *Player, update models.GameUpdate) error {
	opponent := game.Player1
	if sender == opponent {
		opponent = game.Player2
	}

	err := opponent.Conn.WriteJSON(update)
	if err != nil {
		h.logger.Error("Error notifying player", zap.Error(err))
		return err
	}

	return nil
}

func (h *handlers) notifyAll(game *Game, update models.GameUpdate) error {
	if game.Player1 != nil {
		if err := game.Player1.Conn.WriteJSON(update); err != nil {
			h.logger.Error("Error notifying player", zap.Error(err))
			return err
		}
	}

	if game.Player2 != nil {
		if err := game.Player2.Conn.WriteJSON(update); err != nil {
			h.logger.Error("Error notifying player", zap.Error(err))
			return err
		}
	}

	return nil
}

func (h *handlers) GameWSLaunch(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user")
	gameID := r.URL.Query().Get("game")

	if err := UUIDCheck(userID, gameID); err != nil {
		h.logger.Error("Invalid UUID in request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// todo check if we have that game and user

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.logger.Error("Error upgrading connection", zap.Error(err))
		return
	}
	defer conn.Close()

	player := &Player{
		Id:   userID,
		Conn: conn,
	}

	game := h.getOrCreateGame(gameID)

	switch {
	case game.Player1 == nil:
		game.Player1 = player

	case game.Player2 == nil:
		game.Player1 = player
		if err = h.notifyAll(game, models.GameUpdate{}); err != nil {
			h.logger.Error("Error on notifying player", zap.Error(err))
		}

	default:
		h.logger.Error("Error third connection", zap.Any("gameInfo", game), zap.String("thirdUser", player.Id))
		return
	}

	defer func() {
		if game.Player1 == nil && game.Player2 == nil {
			h.removeRoom(game.Id)
		}
	}()

	var action models.GameAction
	for {
		err = conn.ReadJSON(&action)
		if err != nil {
			h.logger.Error("Error reading message", zap.Error(err))
			break
		}

		// todo call game logic, return GameUpdate

		if err = h.notifyOpponent(game, player, models.GameUpdate{}); err != nil {
			h.logger.Error("Error on notifying player", zap.Error(err))
		}
	}
}
