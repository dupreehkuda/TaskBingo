package service

import (
	"context"

	"go.uber.org/zap"
)

// AcceptFriend writes data when user accepts friendship
func (s service) AcceptFriend(ctx context.Context, userID, friendID string) error {
	err := s.repository.AcceptFriend(ctx, userID, friendID)
	if err != nil {
		s.logger.Error("Error occurred in call to repository", zap.Error(err))
		return err
	}

	return nil
}

// DeleteFriend writes data when user cancels/declines friendship
func (s service) DeleteFriend(ctx context.Context, userID, friendID string) error {
	err := s.repository.DeleteFriend(ctx, userID, friendID)
	if err != nil {
		s.logger.Error("Error occurred in call to repository", zap.Error(err))
		return err
	}

	return nil
}

// RequestFriend writes data when user requests friendship
func (s service) RequestFriend(ctx context.Context, userID, friendID string) error {
	err := s.repository.RequestFriend(ctx, userID, friendID)
	if err != nil {
		s.logger.Error("Error occurred in call to repository", zap.Error(err))
		return err
	}

	return nil
}
