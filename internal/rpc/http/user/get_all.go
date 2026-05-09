package user

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

type AllLister interface {
	Run(ctx context.Context) (models.Users, error)
}

type GetAllHandler struct {
	uc     AllLister
	logger *zap.Logger
}

func NewGetAllHandler(uc AllLister, logger *zap.Logger) *GetAllHandler {
	return &GetAllHandler{uc: uc, logger: logger}
}

func (h *GetAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	users, err := h.uc.Run(r.Context())
	if err != nil {
		h.logger.Error("get_all failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(users)
}
