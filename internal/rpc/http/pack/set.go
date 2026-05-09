package pack

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

type Setter interface {
	Run(ctx context.Context, userID string, pack *models.TaskPack) error
}

type SetHandler struct {
	uc     Setter
	logger *zap.Logger
}

func NewSetHandler(uc Setter, logger *zap.Logger) *SetHandler {
	return &SetHandler{uc: uc, logger: logger}
}

func (h *SetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(models.CtxUserIDKey).(string)
	var pack models.TaskPack
	if err := json.NewDecoder(r.Body).Decode(&pack); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.uc.Run(r.Context(), userID, &pack); err != nil {
		h.logger.Error("pack_set failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(pack)
}
