package friend

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/uuidcheck"
)

type Requester interface {
	Run(ctx context.Context, userID, friendID string) error
}

type RequestHandler struct {
	uc     Requester
	logger *zap.Logger
}

func NewRequestHandler(uc Requester, logger *zap.Logger) *RequestHandler {
	return &RequestHandler{uc: uc, logger: logger}
}

func (h *RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	if userID == req.Person {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	if err := h.uc.Run(r.Context(), userID, req.Person); err != nil {
		h.logger.Error("friend_request failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Access-Control-Allow-Credentials", "true")
}
