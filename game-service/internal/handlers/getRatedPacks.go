package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// GetRatedPacks handles getting rated packs operation
func (h handlers) GetRatedPacks(w http.ResponseWriter, r *http.Request) {
	var ctxKey models.LoginKey = "login"
	login := r.Context().Value(ctxKey).(string)

	if login == "" {
		h.logger.Error("Bad login")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.processor.GetRatedPacks()
	if err != nil {
		h.logger.Error("Error in call to processor", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
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
