package repository

import (
	"context"

	"go.uber.org/zap"
)

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
