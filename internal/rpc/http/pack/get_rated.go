package pack

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

type RatedGetter interface {
	Run(ctx context.Context) (models.Packs, error)
}

type GetRatedHandler struct {
	uc     RatedGetter
	logger *zap.Logger
}

func NewGetRatedHandler(uc RatedGetter, logger *zap.Logger) *GetRatedHandler {
	return &GetRatedHandler{uc: uc, logger: logger}
}

func (h *GetRatedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	packs, err := h.uc.Run(r.Context())
	if err != nil {
		h.logger.Error("pack_get_rated failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(packs)
}
