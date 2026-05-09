package game

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/uuidcheck"
)

type Creator interface {
	Run(ctx context.Context, userID, opponentID, packID string) (*models.GameShort, error)
}

type CreateHandler struct {
	uc     Creator
	logger *zap.Logger
}

func NewCreateHandler(uc Creator, logger *zap.Logger) *CreateHandler {
	return &CreateHandler{uc: uc, logger: logger}
}

func (h *CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(models.CtxUserIDKey).(string)
	var req models.NewGameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := uuidcheck.Check(userID, req.OpponentID, req.Pack); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	short, err := h.uc.Run(r.Context(), userID, req.OpponentID, req.Pack)
	if err != nil {
		h.logger.Error("game_create failed", zap.Error(err))
		return
	}
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(short)
}
