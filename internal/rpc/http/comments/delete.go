package comments

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/errs"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/uuidcheck"
)

type Deleter interface {
	Run(ctx context.Context, userID, commentID string) error
}

type DeleteHandler struct {
	uc     Deleter
	logger *zap.Logger
}

func NewDeleteHandler(uc Deleter, logger *zap.Logger) *DeleteHandler {
	return &DeleteHandler{uc: uc, logger: logger}
}

func (h *DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(models.CtxUserIDKey).(string)
	commentID := chi.URLParam(r, "commentID")
	if err := uuidcheck.Check(userID, commentID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := h.uc.Run(r.Context(), userID, commentID)
	if errors.Is(err, errs.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		h.logger.Error("comment_delete failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Access-Control-Allow-Credentials", "true")
}
