package userRepository

import (
	"context"

	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// RequestFriend calls user service to request friendship
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

// AcceptFriend calls user service to accept friendship
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

// DeleteFriend calls user service to delete/cancel friendship
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
