package comments

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/uuidcheck"
)

type Lister interface {
	Run(ctx context.Context, userID, gameID string) ([]models.Comment, error)
}

type ListHandler struct {
	uc     Lister
	logger *zap.Logger
}

func NewListHandler(uc Lister, logger *zap.Logger) *ListHandler {
	return &ListHandler{uc: uc, logger: logger}
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(models.CtxUserIDKey).(string)
	gameID := chi.URLParam(r, "gameID")
	if err := uuidcheck.Check(userID, gameID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	out, err := h.uc.Run(r.Context(), userID, gameID)
	if err != nil {
		h.logger.Error("comment_list failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if out == nil {
		out = []models.Comment{}
	}
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(out)
}
