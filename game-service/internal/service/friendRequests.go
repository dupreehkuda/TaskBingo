package service

import (
	"context"

	"go.uber.org/zap"
)

// RequestFriend calls user repository to request friendship
func (s service) RequestFriend(ctx context.Context, userID, friendID string) error {
	err := s.userRepository.RequestFriend(ctx, userID, friendID)
	if err != nil {
		s.logger.Error("Error occurred in call to user service", zap.Error(err))
		return err
	}

	return nil
}

// AcceptFriend calls user repository to accept friendship
func (s service) AcceptFriend(ctx context.Context, userID, friendID string) error {
	err := s.userRepository.AcceptFriend(ctx, userID, friendID)
	if err != nil {
		s.logger.Error("Error occurred in call to user service", zap.Error(err))
		return err
	}

	return nil
}

// DeleteFriend calls user repository to delete/cancel friendship
func (s service) DeleteFriend(ctx context.Context, userID, friendID string) error {
	err := s.userRepository.DeleteFriend(ctx, userID, friendID)
	if err != nil {
		s.logger.Error("Error occurred in call to user service", zap.Error(err))
		return err
	}

	return nil
}
