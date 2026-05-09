package comments

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/errs"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/uuidcheck"
)

type Editor interface {
	Run(ctx context.Context, userID, commentID, body string) (*models.Comment, error)
}

type EditHandler struct {
	uc     Editor
	logger *zap.Logger
}

func NewEditHandler(uc Editor, logger *zap.Logger) *EditHandler {
	return &EditHandler{uc: uc, logger: logger}
}

func (h *EditHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(models.CtxUserIDKey).(string)
	commentID := chi.URLParam(r, "commentID")
	if err := uuidcheck.Check(userID, commentID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var req models.EditCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Body == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c, err := h.uc.Run(r.Context(), userID, commentID, req.Body)
	if errors.Is(err, errs.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		h.logger.Error("comment_edit failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(c)
}
