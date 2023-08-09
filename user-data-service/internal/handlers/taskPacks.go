package handlers

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

// SetNewTaskPack handles the operation of setting & assigning new pack
func (h *Handlers) SetNewTaskPack(ctx context.Context, req *api.NewTaskPackRequest) (*api.Empty, error) {
	pack := &models.TaskPack{
		ID: req.PackID,
		Pack: models.Pack{
			Title: req.Pack.Title,
			Tasks: req.Pack.Tasks,
		},
	}

	if err := h.service.AddTaskPack(ctx, req.UserID, pack); err != nil {
		h.logger.Error("Unable to call service", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}

// GetTaskPacks handles the operation of getting a pack
func (h *Handlers) GetTaskPacks(ctx context.Context, req *api.TaskPacksRequest) (*api.TaskPacksResponse, error) {
	packs, err := h.service.GetTaskPacks(ctx, req.Ids...)
	if err != nil {
		h.logger.Error("Unable to call service", zap.Error(err))
		return nil, err
	}

	return mapToMultiplePacks(packs), nil
}

// GetRatedPacks handles the operation of getting some packs in desc rating order
func (h *Handlers) GetRatedPacks(ctx context.Context, _ *api.Empty) (*api.GetMultiplePacksResponse, error) {
	resp, err := h.service.GetRatedPacks(ctx)
	if err != nil {
		h.logger.Error("Unable to call service", zap.Error(err))
		return nil, err
	}

	return &api.GetMultiplePacksResponse{Packs: mapToPacks(resp)}, nil
}

// LikePack handles the operation of user liking the pack
func (h *Handlers) LikePack(ctx context.Context, req *api.LikeOrRatePackRequest) (*api.Empty, error) {
	err := h.service.LikePack(ctx, req.UserID, req.Pack, 1)
	if err != nil {
		h.logger.Error("Error occurred in call to processor", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}

// DislikePack handles the operation of user disliking the pack
func (h *Handlers) DislikePack(ctx context.Context, req *api.LikeOrRatePackRequest) (*api.Empty, error) {
	err := h.service.LikePack(ctx, req.UserID, req.Pack, -1)
	if err != nil {
		h.logger.Error("Error occurred in call to processor", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}

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
