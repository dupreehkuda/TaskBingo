package solo

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/uuidcheck"
)

type Starter interface {
	Run(ctx context.Context, userID, packID string) (*models.SoloStartResponse, error)
}

type StartHandler struct {
	uc     Starter
	logger *zap.Logger
}

func NewStartHandler(uc Starter, logger *zap.Logger) *StartHandler {
	return &StartHandler{uc: uc, logger: logger}
}

func (h *StartHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(models.CtxUserIDKey).(string)
	var req models.SoloStartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := uuidcheck.Check(userID, req.PackID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := h.uc.Run(r.Context(), userID, req.PackID)
	if err != nil {
		h.logger.Error("solo start failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
