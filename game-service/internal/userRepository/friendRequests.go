package userRepository

import (
	"context"

	"go.uber.org/zap"
)

// RequestFriend calls user service to request friendship
func (u userRepository) RequestFriend(ctx context.Context, userID, friendID string) error {
	_, err := u.conn.RequestFriend(ctx, mapToFriendRequest(userID, friendID))

	if err != nil {
		u.logger.Error("Error occurred on user service", zap.Error(err))
		return err
	}

	return nil
}

// AcceptFriend calls user service to accept friendship
func (u userRepository) AcceptFriend(ctx context.Context, userID, friendID string) error {
	_, err := u.conn.AcceptFriend(ctx, mapToFriendRequest(userID, friendID))

	if err != nil {
		u.logger.Error("Error occurred on user service", zap.Error(err))
		return err
	}

	return nil
}

// DeleteFriend calls user service to delete/cancel friendship
func (u userRepository) DeleteFriend(ctx context.Context, userID, friendID string) error {
	_, err := u.conn.DeleteFriend(ctx, mapToFriendRequest(userID, friendID))

	if err != nil {
		u.logger.Error("Error occurred on user service", zap.Error(err))
		return err
	}

	return nil
}
