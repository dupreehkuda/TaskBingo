package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	"go.uber.org/zap"
)

// CreateGame handles the operation of creating new game
func (h handlers) CreateGame(w http.ResponseWriter, r *http.Request) {
	var ctxKey models.LoginKey = "login"
	login := r.Context().Value(ctxKey).(string)

	if login == "" {
		h.logger.Error("Bad login")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req models.NewGameRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		h.logger.Error("Unable to decode JSON", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if req.User == "" || req.Pack == "" {
		h.logger.Info("Request empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.processor.CreateGame(login, req.User, req.Pack)
	if err != nil {
		h.logger.Error("Error getting data", zap.Error(err))
		return
	}

	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Origin", "*")
}
