package pack

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

type Rater interface {
	Run(ctx context.Context, userID, packID string) error
}

type RateHandler struct {
	uc     Rater
	logger *zap.Logger
}

func NewRateHandler(uc Rater, logger *zap.Logger) *RateHandler {
	return &RateHandler{uc: uc, logger: logger}
}

func (h *RateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(models.CtxUserIDKey).(string)
	var req models.PackAction
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.uc.Run(r.Context(), userID, req.Pack); err != nil {
		h.logger.Error("pack_rate failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Access-Control-Allow-Credentials", "true")
}
