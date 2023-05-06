package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// RateTaskPack handles pack star operations
func (h *handlers) RateTaskPack(w http.ResponseWriter, r *http.Request) {
	var ctxKey models.UserIDKey = "userID"
	userID := r.Context().Value(ctxKey).(string)

	var req models.TaskPackRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		h.logger.Error("Unable to decode JSON", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = UUIDCheck(userID, req.TaskID); err != nil {
		h.logger.Error("Invalid UUID in request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.RateTaskPack(userID, req.TaskID)
	if err != nil {
		h.logger.Error("Error getting data", zap.Error(err))
		return
	}

	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Origin", "*")
}

// UnrateTaskPack handles pack unstar operations
func (h *handlers) UnrateTaskPack(w http.ResponseWriter, r *http.Request) {
	var ctxKey models.UserIDKey = "userID"
	userID := r.Context().Value(ctxKey).(string)

	var req models.TaskPackRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		h.logger.Error("Unable to decode JSON", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = UUIDCheck(userID, req.TaskID); err != nil {
		h.logger.Error("Invalid UUID in request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.UnrateTaskPack(userID, req.TaskID)
	if err != nil {
		h.logger.Error("Error getting data", zap.Error(err))
		return
	}

	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Origin", "https://taskbingo.com")
}
