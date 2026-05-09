package pack

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

type ManyGetter interface {
	Run(ctx context.Context, requesterID string, packIDs []string) (models.Packs, error)
}

type GetManyHandler struct {
	uc     ManyGetter
	logger *zap.Logger
}

func NewGetManyHandler(uc ManyGetter, logger *zap.Logger) *GetManyHandler {
	return &GetManyHandler{uc: uc, logger: logger}
}

func (h *GetManyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(models.CtxUserIDKey).(string)
	var req models.TaskPacksRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	packs, err := h.uc.Run(r.Context(), userID, req.PackIDs)
	if err != nil {
		h.logger.Error("pack_get_many failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(packs)
}
