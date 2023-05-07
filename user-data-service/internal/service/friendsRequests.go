package service

import (
	"context"

	"go.uber.org/zap"
)

func (s service) AcceptFriend(ctx context.Context, userID, friendID string) error {
	err := s.repository.AcceptFriend(ctx, userID, friendID)
	if err != nil {
		s.logger.Error("Error occurred in call to repository", zap.Error(err))
		return err
	}

	return nil
}

func (s service) DeleteFriend(ctx context.Context, userID, friendID string) error {
	err := s.repository.DeleteFriend(ctx, userID, friendID)
	if err != nil {
		s.logger.Error("Error occurred in call to repository", zap.Error(err))
		return err
	}

	return nil
}

func (s service) RequestFriend(ctx context.Context, userID, friendID string) error {
	err := s.repository.RequestFriend(ctx, userID, friendID)
	if err != nil {
		s.logger.Error("Error occurred in call to repository", zap.Error(err))
		return err
	}

	return nil
}
