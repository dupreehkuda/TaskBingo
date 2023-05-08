package service

import (
	"context"

	"go.uber.org/zap"
)

// RateTaskPack adds to pack rating
func (s service) RateTaskPack(ctx context.Context, userID, pack string) error {
	err := s.userRepository.RateTaskPack(ctx, userID, pack)
	if err != nil {
		s.logger.Error("Error occurred calling user repository", zap.Error(err))
		return err
	}

	return nil
}

// UnrateTaskPack removes from pack rating
func (s service) UnrateTaskPack(ctx context.Context, userID, pack string) error {
	err := s.userRepository.UnrateTaskPack(ctx, userID, pack)
	if err != nil {
		s.logger.Error("Error occurred calling user repository", zap.Error(err))
		return err
	}

	return nil
}
