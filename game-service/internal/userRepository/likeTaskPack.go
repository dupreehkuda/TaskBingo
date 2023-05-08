package userRepository

import (
	"context"

	"go.uber.org/zap"
)

// LikeTaskPack calls user service to like pack by user
func (u userRepository) LikeTaskPack(ctx context.Context, userID, pack string) error {
	_, err := u.conn.LikePack(ctx, mapToLikeOrRate(userID, pack))

	if err != nil {
		u.logger.Error("Error when liking pack", zap.Error(err))
		return err
	}

	return nil
}

// DislikeTaskPack calls user service to dislike pack by user
func (u userRepository) DislikeTaskPack(ctx context.Context, userID, pack string) error {
	_, err := u.conn.DislikePack(ctx, mapToLikeOrRate(userID, pack))

	if err != nil {
		u.logger.Error("Error when disliking pack", zap.Error(err))
		return err
	}

	return nil
}
