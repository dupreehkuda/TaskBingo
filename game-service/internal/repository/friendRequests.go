package repository

import (
	"context"

	"go.uber.org/zap"
)

// RequestFriend calls user service to request friendship
func (r repository) RequestFriend(ctx context.Context, userID, friendID string) error {
	_, err := r.conn.RequestFriend(ctx, mapToFriendRequest(userID, friendID))

	if err != nil {
		r.logger.Error("Error occurred on user service", zap.Error(err))
		return err
	}

	return nil
}

// AcceptFriend calls user service to accept friendship
func (r repository) AcceptFriend(ctx context.Context, userID, friendID string) error {
	_, err := r.conn.AcceptFriend(ctx, mapToFriendRequest(userID, friendID))

	if err != nil {
		r.logger.Error("Error occurred on user service", zap.Error(err))
		return err
	}

	return nil
}

// DeleteFriend calls user service to delete/cancel friendship
func (r repository) DeleteFriend(ctx context.Context, userID, friendID string) error {
	_, err := r.conn.DeleteFriend(ctx, mapToFriendRequest(userID, friendID))

	if err != nil {
		r.logger.Error("Error occurred on user service", zap.Error(err))
		return err
	}

	return nil
}
