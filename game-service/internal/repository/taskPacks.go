package repository

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

func (r repository) SetNewTaskPack(ctx context.Context, userID string, pack *models.TaskPack) error {
	if _, err := r.conn.SetNewTaskPack(ctx, mapToNewTaskPack(userID, pack)); err != nil {
		r.logger.Error("Error setting pack", zap.Error(err))
		return err
	}

	return nil
}

func (r repository) GetTaskPacks(ctx context.Context, packIDs ...string) (*models.Packs, error) {
	packs, err := r.conn.GetTaskPacks(ctx, &api.TaskPacksRequest{Ids: packIDs})
	if err != nil {
		r.logger.Error("Error when getting user data", zap.Error(err))
		return nil, err
	}

	return mapFromTaskPack(packs), nil
}

// GetRatedPacks retrieves most rated packs from user service
func (r repository) GetRatedPacks(ctx context.Context) (*models.Packs, error) {
	resp, err := r.conn.GetRatedPacks(ctx, &api.Empty{})
	if err != nil {
		r.logger.Error("Error when getting user data", zap.Error(err))
		return nil, err
	}

	if len(resp.Packs) == 0 {
		return &models.Packs{}, nil
	}

	return mapFromRatedPacks(resp), nil
}

// LikeTaskPack calls user service to like pack by user
func (r repository) LikeTaskPack(ctx context.Context, userID, pack string) error {
	_, err := r.conn.LikePack(ctx, mapToLikeOrRate(userID, pack))

	if err != nil {
		r.logger.Error("Error when liking pack", zap.Error(err))
		return err
	}

	return nil
}

// DislikeTaskPack calls user service to dislike pack by user
func (r repository) DislikeTaskPack(ctx context.Context, userID, pack string) error {
	_, err := r.conn.DislikePack(ctx, mapToLikeOrRate(userID, pack))

	if err != nil {
		r.logger.Error("Error when disliking pack", zap.Error(err))
		return err
	}

	return nil
}

// RateTaskPack calls user service to add rating to the pack
func (r repository) RateTaskPack(ctx context.Context, userID, pack string) error {
	_, err := r.conn.RatePack(ctx, mapToLikeOrRate(userID, pack))

	if err != nil {
		r.logger.Error("Error when staring pack", zap.Error(err))
		return err
	}

	return nil
}

// UnrateTaskPack calls user service to remove rating from the pack
func (r repository) UnrateTaskPack(ctx context.Context, userID, pack string) error {
	_, err := r.conn.UnratePack(ctx, mapToLikeOrRate(userID, pack))

	if err != nil {
		r.logger.Error("Error when unstaring pack", zap.Error(err))
		return err
	}

	return nil
}
