package handlers

import (
	"net/http"

	"github.com/mailru/easyjson"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// GetRatedPacks handles getting rated packs operation
func (h *handlers) GetRatedPacks(w http.ResponseWriter, r *http.Request) {
	var ctxKey models.UserIDKey = "userID"
	userID := r.Context().Value(ctxKey).(string)

	if err := UUIDCheck(userID); err != nil {
		h.logger.Error("Invalid UUID in request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.service.GetRatedPacks(r.Context())
	if err != nil {
		h.logger.Error("Error in call to processor", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resultJSON, err := easyjson.Marshal(resp)
	if err != nil {
		h.logger.Error("Error marshaling data", zap.Error(err))
		return
	}

	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(resultJSON)
	if err != nil {
		h.logger.Error("Unable to write response", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
