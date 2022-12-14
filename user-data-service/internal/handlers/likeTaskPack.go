package handlers

import (
	"context"

	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

// LikePack handles the operation of user liking the pack
func (h *Handlers) LikePack(ctx context.Context, req *api.LikePackRequest) (*api.Empty, error) {
	err := h.processor.LikePack(req.Login, req.Pack, 1)
	if err != nil {
		h.logger.Error("Error occurred in call to processor", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}

// DislikePack handles the operation of user disliking the pack
func (h *Handlers) DislikePack(ctx context.Context, req *api.LikePackRequest) (*api.Empty, error) {
	err := h.processor.LikePack(req.Login, req.Pack, -1)
	if err != nil {
		h.logger.Error("Error occurred in call to processor", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}
