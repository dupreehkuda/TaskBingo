package handlers

import (
	"context"

	"go.uber.org/zap"

	i "github.com/dupreehkuda/TaskBingo/user-data-service/internal/interfaces"
	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

type Handlers struct {
	storage i.Stored
	logger  *zap.Logger
}

// New creates new instance of handlers
func New(storage i.Stored, logger *zap.Logger) *Handlers {
	return &Handlers{
		storage: storage,
		logger:  logger,
	}
}

func (h *Handlers) Get(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	h.logger.Debug("holy shit it works")
	resp, err := h.storage.Ping(req.Id)
	if err != nil {
		h.logger.Error("Unable to call storage", zap.Error(err))
		return nil, err
	}

	return &api.GetResponse{
		Nickname: resp.UserID,
		Points:   int32(resp.Points),
		Email:    resp.Email,
	}, nil
}
