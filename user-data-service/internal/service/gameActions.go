package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// GetGame creates new game instance in service
func (s service) GetGame(ctx context.Context, gameID string) (*models.Game, error) {
	game, err := s.repository.GetGame(ctx, gameID)
	if err != nil {
		s.logger.Error("Error occurred in call to repository", zap.Error(err))
		return nil, err
	}

	return game, nil
}

// AchieveGame creates new game instance in service
func (s service) AchieveGame(ctx context.Context, game *models.Game) error {
	if err := s.repository.AchieveGame(ctx, game); err != nil {
		s.logger.Error("Error occurred in call to repository", zap.Error(err))
		return err
	}

	return nil
}
