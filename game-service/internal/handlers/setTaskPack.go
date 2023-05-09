package handlers

import (
	"io"
	"net/http"

	"github.com/mailru/easyjson"
	"go.uber.org/zap"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// SetTaskPack handles addition of new task packs
func (h *handlers) SetTaskPack(w http.ResponseWriter, r *http.Request) {
	var ctxKey models.UserIDKey = "userID"
	userID := r.Context().Value(ctxKey).(string)

	var req models.TaskPack

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

	if req.Pack.Title == "" {
		h.logger.Info("Pack title empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = UUIDCheck(userID); err != nil {
		h.logger.Error("Invalid UUID in request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.SetTaskPack(r.Context(), userID, &req)

	switch {
	case err == errs.ErrPackAlreadyExists:
		w.WriteHeader(http.StatusConflict)
		return
	case err != nil:
		h.logger.Error("Error occurred in call to processor", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Origin", "*")
}
