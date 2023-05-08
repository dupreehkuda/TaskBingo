package handlers

import (
	"net/http"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

var wsUpgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// getOrCreateRoom retrieves game from the hub or repository
func (h *handlers) getOrCreateRoom(gameID string) (*models.Room, error) {
	h.hub.Mu.Lock()
	defer h.hub.Mu.Unlock()

	if game, ok := h.hub.Rooms[gameID]; ok {
		return game, nil
	}

	game, err := h.service.GetRoom(gameID)
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
	var ctxKey models.UserIDKey = "userID"
	userID := r.Context().Value(ctxKey).(string)
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

	room, err := h.getOrCreateRoom(gameID)
	if err != nil {
		h.logger.Error("Error getting game", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch {
	case room.Player1 == nil:
		room.Player1 = player

	case room.Player2 == nil:
		room.Player2 = player

	default:
		h.logger.Error("Error third connection", zap.Any("gameInfo", room), zap.String("thirdUser", player.Id))
		return
	}

	defer func() {
		if room.Player1 == nil && room.Player2 == nil {
			h.removeRoom(room.Id)
		}
	}()

	var action models.GameAction
	for {
		err = conn.ReadJSON(&action)
		if err != nil {
			if err == websocket.ErrCloseSent {
				break
			}

			h.logger.Debug("Error reading message", zap.Error(err))
			break
		}

		if err = UUIDCheck(userID); err != nil || len(action.Numbers) != 16 {
			h.logger.Error("Invalid UUID or missing numbers", zap.Error(err))
			continue
		}

		update, err := h.service.UpdateGame(room, &action)
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
