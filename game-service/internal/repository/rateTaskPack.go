package repository

import (
	"context"

	"go.uber.org/zap"
)

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
