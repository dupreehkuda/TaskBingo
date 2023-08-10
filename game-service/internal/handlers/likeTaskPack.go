package handlers

import (
	"io"
	"net/http"

	"github.com/mailru/easyjson"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// LikeTaskPack handles user's like operations
func (h *handlers) LikeTaskPack(w http.ResponseWriter, r *http.Request) {
	var ctxKey models.UserIDKey = "userID"
	userID := r.Context().Value(ctxKey).(string)

	var req models.TaskPackRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Error("Unable to read body", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = easyjson.Unmarshal(body, &req); err != nil {
		h.logger.Error("Unable to decode JSON", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = UUIDCheck(userID, req.TaskID); err != nil {
		h.logger.Error("Invalid UUID in request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.LikeTaskPack(r.Context(), userID, req.TaskID)
	if err != nil {
		h.logger.Error("Error getting data", zap.Error(err))
		return
	}

	w.Header().Add("Access-Control-Allow-Credentials", "true")
}

// DislikeTaskPack handles user's dislike operations
func (h *handlers) DislikeTaskPack(w http.ResponseWriter, r *http.Request) {
	var ctxKey models.UserIDKey = "userID"
	userID := r.Context().Value(ctxKey).(string)

	var req models.TaskPackRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Error("Unable to read body", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = easyjson.Unmarshal(body, &req); err != nil {
		h.logger.Error("Unable to decode JSON", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = UUIDCheck(userID, req.TaskID); err != nil {
		h.logger.Error("Invalid UUID in request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.DislikeTaskPack(r.Context(), userID, req.TaskID)
	if err != nil {
		h.logger.Error("Error getting data", zap.Error(err))
		return
	}

	w.Header().Add("Access-Control-Allow-Credentials", "true")
}
