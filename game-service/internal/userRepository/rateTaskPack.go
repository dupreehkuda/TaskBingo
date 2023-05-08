package userRepository

import (
	"context"

	"go.uber.org/zap"
)

// RateTaskPack calls user service to add rating to the pack
func (u userRepository) RateTaskPack(ctx context.Context, userID, pack string) error {
	_, err := u.conn.RatePack(ctx, mapToLikeOrRate(userID, pack))

	if err != nil {
		u.logger.Error("Error when staring pack", zap.Error(err))
		return err
	}

	return nil
}

// UnrateTaskPack calls user service to remove rating from the pack
func (u userRepository) UnrateTaskPack(ctx context.Context, userID, pack string) error {
	_, err := u.conn.UnratePack(ctx, mapToLikeOrRate(userID, pack))

	if err != nil {
		u.logger.Error("Error when unstaring pack", zap.Error(err))
		return err
	}

	return nil
}
