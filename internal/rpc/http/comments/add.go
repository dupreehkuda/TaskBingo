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

type Adder interface {
	Run(ctx context.Context, userID, gameID, body string) (*models.Comment, error)
}

type AddHandler struct {
	uc     Adder
	logger *zap.Logger
}

func NewAddHandler(uc Adder, logger *zap.Logger) *AddHandler {
	return &AddHandler{uc: uc, logger: logger}
}

func (h *AddHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(models.CtxUserIDKey).(string)
	gameID := chi.URLParam(r, "gameID")
	if err := uuidcheck.Check(userID, gameID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var req models.AddCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Body == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c, err := h.uc.Run(r.Context(), userID, gameID, req.Body)
	if err != nil {
		h.logger.Error("comment_add failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(c)
}
