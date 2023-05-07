package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// CreateGame creates new game instance in service
func (s service) CreateGame(ctx context.Context, game *models.Game) error {
	if err := s.repository.CreateGame(ctx, game); err != nil {
		s.logger.Error("Error occurred in call to repository", zap.Error(err))
		return err
	}

	return nil
}

func (s service) AcceptGame(ctx context.Context, userID, gameID string) error {
	if err := s.repository.AcceptGame(ctx, userID, gameID); err != nil {
		s.logger.Error("Error occurred in call to repository", zap.Error(err))
		return err
	}

	return nil
}

func (s service) DeleteGame(ctx context.Context, userID, gameID string) error {
	if err := s.repository.DeleteGame(ctx, userID, gameID); err != nil {
		s.logger.Error("Error occurred in call to repository", zap.Error(err))
		return err
	}

	return nil
}
