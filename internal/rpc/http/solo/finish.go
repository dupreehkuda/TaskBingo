package solo

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/errs"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/uuidcheck"
)

type Finisher interface {
	Run(ctx context.Context, userID, gameID string, userNumbers []int32) (*models.SoloFinishResponse, error)
}

type FinishHandler struct {
	uc     Finisher
	logger *zap.Logger
}

func NewFinishHandler(uc Finisher, logger *zap.Logger) *FinishHandler {
	return &FinishHandler{uc: uc, logger: logger}
}

func (h *FinishHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(models.CtxUserIDKey).(string)
	var req models.SoloFinishRequest
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
	resp, err := h.uc.Run(r.Context(), userID, req.GameID, req.UserNumbers)
	if errors.Is(err, errs.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		h.logger.Error("solo finish failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
