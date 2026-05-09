package user

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

type DataGetter interface {
	Run(ctx context.Context, userID string) (*models.UserAccountInfo, error)
}

type GetDataHandler struct {
	uc     DataGetter
	logger *zap.Logger
}

func NewGetDataHandler(uc DataGetter, logger *zap.Logger) *GetDataHandler {
	return &GetDataHandler{uc: uc, logger: logger}
}

func (h *GetDataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(models.CtxUserIDKey).(string)
	data, err := h.uc.Run(r.Context(), userID)
	if err != nil {
		h.logger.Error("get_data failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}
