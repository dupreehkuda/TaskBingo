package handlers

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

// LikePack handles the operation of user liking the pack
func (h *Handlers) LikePack(ctx context.Context, req *api.LikeOrRatePackRequest) (*api.Empty, error) {
	userID, err := uuid.Parse(req.UserID.Id)
	if err != nil {
		h.logger.Error("Unable to parse uuid", zap.Error(err))
		return nil, err
	}

	packID, err := uuid.Parse(req.Pack)
	if err != nil {
		h.logger.Error("Unable to parse uuid", zap.Error(err))
		return nil, err
	}

	err = h.processor.LikePack(userID, packID, 1)
	if err != nil {
		h.logger.Error("Error occurred in call to processor", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}

// DislikePack handles the operation of user disliking the pack
func (h *Handlers) DislikePack(ctx context.Context, req *api.LikeOrRatePackRequest) (*api.Empty, error) {
	userID, err := uuid.Parse(req.UserID.Id)
	if err != nil {
		h.logger.Error("Unable to parse uuid", zap.Error(err))
		return nil, err
	}

	packID, err := uuid.Parse(req.Pack)
	if err != nil {
		h.logger.Error("Unable to parse uuid", zap.Error(err))
		return nil, err
	}

	err = h.processor.LikePack(userID, packID, -1)
	if err != nil {
		h.logger.Error("Error occurred in call to processor", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}
