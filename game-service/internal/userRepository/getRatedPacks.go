package userRepository

import (
	"context"

	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// GetRatedPacks retrieves most rated packs from user service
func (u userRepository) GetRatedPacks(ctx context.Context) ([]string, error) {
	resp, err := u.conn.GetRatedPacks(ctx, &api.Empty{})
	if err != nil {
		u.logger.Error("Error when getting user data", zap.Error(err))
		return nil, err
	}

	return mapFromRatedPacks(resp), nil
}
