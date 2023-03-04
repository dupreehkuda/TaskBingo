package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// GetTaskPack handles getting one task pack operations
func (h handlers) GetTaskPack(w http.ResponseWriter, r *http.Request) {
	var ctxKey models.UserIDKey = "userID"
	userID := r.Context().Value(ctxKey).(string)

	if userID == "" {
		h.logger.Error("Bad login")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req models.TaskPackRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		h.logger.Error("Unable to decode JSON", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if req.TaskID == "" {
		h.logger.Info("Request empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.processor.GetTaskPack(req.TaskID)

	switch {
	case err == errs.ErrNoSuchPack:
		w.WriteHeader(http.StatusNoContent)
		return
	case err != nil:
		h.logger.Error("Unable to return task pack", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resultJSON, err := json.Marshal(resp)
	if err != nil {
		h.logger.Error("Error marshaling data", zap.Error(err))
		return
	}

	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	_, err = w.Write(resultJSON)
	if err != nil {
		h.logger.Error("Unable to write response", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
