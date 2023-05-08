package service

import (
	"context"

	"go.uber.org/zap"
)

// LikeTaskPack likes pack by user
func (s service) LikeTaskPack(ctx context.Context, userID, pack string) error {
	err := s.userRepository.LikeTaskPack(ctx, userID, pack)
	if err != nil {
		s.logger.Error("Error occurred calling user repository", zap.Error(err))
		return err
	}

	return nil
}

// DislikeTaskPack dislikes pack by user
func (s service) DislikeTaskPack(ctx context.Context, userID, pack string) error {
	err := s.userRepository.DislikeTaskPack(ctx, userID, pack)
	if err != nil {
		s.logger.Error("Error occurred calling user repository", zap.Error(err))
		return err
	}

	return nil
}
