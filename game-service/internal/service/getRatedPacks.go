package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// GetRatedPacks gets some most rated packs
func (s service) GetRatedPacks(ctx context.Context) (*models.Packs, error) {
	rated, err := s.userRepository.GetRatedPacks(ctx)
	if err != nil {
		s.logger.Error("Error occurred in call to user repository", zap.Error(err))
		return nil, err
	}

	if len(rated) == 0 {
		return &models.Packs{}, nil
	}

	packs, err := s.taskRepository.GetMultiplePacks(ctx, rated)
	if err != nil {
		s.logger.Error("Error occurred in call to task repository", zap.Error(err))
		return nil, err
	}

	return packs, nil
}
