package userClient

import (
	"context"

	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// RateTaskPack calls user service to add rating to the pack
func (u userClient) RateTaskPack(userID, pack string) error {
	_, err := u.conn.RatePack(context.Background(), &api.LikeOrRatePackRequest{
		UserID: userID,
		Pack:   pack,
	})

	if err != nil {
		u.logger.Error("Error when staring pack", zap.Error(err))
		return err
	}

	return nil
}

// UnrateTaskPack calls user service to remove rating from the pack
func (u userClient) UnrateTaskPack(userID, pack string) error {
	_, err := u.conn.UnratePack(context.Background(), &api.LikeOrRatePackRequest{
		UserID: userID,
		Pack:   pack,
	})

	if err != nil {
		u.logger.Error("Error when unstaring pack", zap.Error(err))
		return err
	}

	return nil
}
