package solo

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/uuidcheck"
)

type Progresser interface {
	Run(ctx context.Context, userID, gameID string, userNumbers []int32) error
}

type ProgressHandler struct {
	uc     Progresser
	logger *zap.Logger
}

func NewProgressHandler(uc Progresser, logger *zap.Logger) *ProgressHandler {
	return &ProgressHandler{uc: uc, logger: logger}
}

func (h *ProgressHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(models.CtxUserIDKey).(string)
	var req models.SoloProgressRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := uuidcheck.Check(userID, req.GameID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(req.UserNumbers) != 16 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.uc.Run(r.Context(), userID, req.GameID, req.UserNumbers); err != nil {
		h.logger.Error("solo progress failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Access-Control-Allow-Credentials", "true")
}
