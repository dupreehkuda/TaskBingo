package handlers

import (
	"context"

	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

// RatePack handles the operation of rating a pack by user
func (h *Handlers) RatePack(ctx context.Context, req *api.LikeOrRatePackRequest) (*api.Empty, error) {
	err := h.service.RatePack(ctx, req.UserID, req.Pack, 1)
	if err != nil {
		h.logger.Error("Error occurred in call to processor", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}

// UnratePack handles the operation of unrating a pack by user
func (h *Handlers) UnratePack(ctx context.Context, req *api.LikeOrRatePackRequest) (*api.Empty, error) {
	err := h.service.RatePack(ctx, req.UserID, req.Pack, -1)
	if err != nil {
		h.logger.Error("Error occurred in call to processor", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}
