package friend

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/uuidcheck"
)

type Deleter interface {
	Run(ctx context.Context, userID, friendID string) error
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
	var req models.FriendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := uuidcheck.Check(userID, req.Person); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.uc.Run(r.Context(), userID, req.Person); err != nil {
		h.logger.Error("friend_delete failed", zap.Error(err))
		return
	}
	w.Header().Add("Access-Control-Allow-Credentials", "true")
}
