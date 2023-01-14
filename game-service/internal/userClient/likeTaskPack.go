package userClient

import (
	"context"

	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// LikeTaskPack calls user service to like pack by user
func (u userClient) LikeTaskPack(login, pack string) error {
	_, err := u.conn.LikePack(context.Background(), &api.LikeOrRatePackRequest{
		Login: login,
		Pack:  pack,
	})

	if err != nil {
		u.logger.Error("Error when liking pack", zap.Error(err))
		return err
	}

	return nil
}

// DislikeTaskPack calls user service to dislike pack by user
func (u userClient) DislikeTaskPack(login, pack string) error {
	_, err := u.conn.DislikePack(context.Background(), &api.LikeOrRatePackRequest{
		Login: login,
		Pack:  pack,
	})

	if err != nil {
		u.logger.Error("Error when disliking pack", zap.Error(err))
		return err
	}

	return nil
}
