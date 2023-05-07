package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

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

func (h *handlers) removeRoom(id string) {
	h.hub.Mu.Lock()
	defer h.hub.Mu.Unlock()

	if _, ok := h.hub.Rooms[id]; !ok {
		return
	}

	delete(h.hub.Rooms, id)
}

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

func (h *handlers) GameWSLaunch(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user")
	gameID := r.URL.Query().Get("game")

	h.logger.Debug("ids", zap.String("userID", userID), zap.String("gameID", gameID))

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

		// todo send normal update
		if err = h.notifyAll(room, &models.GameUpdate{
			Status:  666,
			UserID:  "hi",
			Numbers: []int32{0},
		}); err != nil {
			h.logger.Error("Error on notifying player", zap.Error(err))
		}

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

		h.logger.Debug("new action", zap.Any("val", action))
		if err = UUIDCheck(userID); err != nil || len(action.Numbers) != 16 {
			h.logger.Error("Invalid UUID or missing numbers", zap.Error(err))
			continue
		}

		fmt.Println("before calling func", &room)
		update, err := h.service.UpdateGame(room, &action)
		if err != nil {
			h.logger.Error("Error getting update", zap.Error(err))
		}

		h.logger.Debug("status", zap.Any("game", room))
		fmt.Println("after calling func", &room)
		if update == nil {
			continue
		}

		// todo improve check. 6 shows game ending flag
		if update.Status == 6 {
			if err = h.notifyAll(room, update); err != nil {
				h.logger.Error("Error on notifying player", zap.Error(err))
			}
			return
		}

		if err = h.notifyOpponent(room, player, update); err != nil {
			h.logger.Error("Error on notifying player", zap.Error(err))
		}
	}

	h.logger.Debug("Connection closed", zap.String("userID", userID))
}
