package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// CreateGame handles the operation of creating new game
func (h handlers) CreateGame(w http.ResponseWriter, r *http.Request) {
	var ctxKey models.UserIDKey = "userID"
	userID := r.Context().Value(ctxKey).(string)

	var req models.NewGameRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		h.logger.Error("Unable to decode JSON", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = UUIDCheck(userID, req.OpponentID, req.Pack); err != nil {
		h.logger.Error("Invalid UUID in request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.CreateGame(userID, req.OpponentID, req.Pack)
	if err != nil {
		h.logger.Error("Error getting data", zap.Error(err))
		return
	}

	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Origin", "*")
}
