package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// SetTaskPack handles addition of setting new task packs
func (h handlers) SetTaskPack(w http.ResponseWriter, r *http.Request) {
	var ctxKey models.LoginKey = "login"
	login := r.Context().Value(ctxKey).(string)

	if login == "" {
		h.logger.Error("Bad login")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req := &models.TaskPack{}

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

	err = h.processor.SetTaskPack(login, req)

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
