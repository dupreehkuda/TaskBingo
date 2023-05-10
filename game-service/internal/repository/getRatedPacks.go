package repository

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

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
