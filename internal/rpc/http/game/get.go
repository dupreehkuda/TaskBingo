package game

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/uuidcheck"
)

type Getter interface {
	Run(ctx context.Context, gameID string) (*models.Game, error)
}

type GetHandler struct {
	uc     Getter
	logger *zap.Logger
}

func NewGetHandler(uc Getter, logger *zap.Logger) *GetHandler {
	return &GetHandler{uc: uc, logger: logger}
}

func (h *GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	g, err := h.uc.Run(r.Context(), req.GameID)
	if err != nil {
		h.logger.Error("game_get failed", zap.Error(err))
		return
	}
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(g)
}
