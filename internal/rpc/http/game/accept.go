package game

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/uuidcheck"
)

type Acceptor interface {
	Run(ctx context.Context, userID, gameID string) error
}

type AcceptHandler struct {
	uc     Acceptor
	logger *zap.Logger
}

func NewAcceptHandler(uc Acceptor, logger *zap.Logger) *AcceptHandler {
	return &AcceptHandler{uc: uc, logger: logger}
}

func (h *AcceptHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(models.CtxUserIDKey).(string)
	var req models.StatusGameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := uuidcheck.Check(userID, req.GameID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.uc.Run(r.Context(), userID, req.GameID); err != nil {
		h.logger.Error("game_accept failed", zap.Error(err))
		return
	}
	w.Header().Add("Access-Control-Allow-Credentials", "true")
}
