package userRepository

import (
	"context"

	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

func (u userRepository) RequestFriend(userID, friendID string) error {
	_, err := u.conn.RequestFriend(context.Background(), &api.FriendRequest{
		UserID:   userID,
		FriendID: friendID,
	})

	if err != nil {
		u.logger.Error("Error occurred on user service", zap.Error(err))
		return err
	}

	return nil
}

func (u userRepository) AcceptFriend(userID, friendID string) error {
	_, err := u.conn.AcceptFriend(context.Background(), &api.FriendRequest{
		UserID:   userID,
		FriendID: friendID,
	})

	if err != nil {
		u.logger.Error("Error occurred on user service", zap.Error(err))
		return err
	}

	return nil
}

func (u userRepository) DeleteFriend(userID, friendID string) error {
	_, err := u.conn.DeleteFriend(context.Background(), &api.FriendRequest{
		UserID:   userID,
		FriendID: friendID,
	})

	if err != nil {
		u.logger.Error("Error occurred on user service", zap.Error(err))
		return err
	}

	return nil
}
