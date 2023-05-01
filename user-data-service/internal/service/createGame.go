package service

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// CreateGame creates new game instance in service
func (s service) CreateGame(game *models.Game) error {
	if err := s.repository.CreateGame(game); err != nil {
		s.logger.Error("Error occurred in call to repository", zap.Error(err))
		return err
	}

	return nil
}

func (s service) AcceptGame(userID, gameID string) error {
	if err := s.repository.AcceptGame(userID, gameID); err != nil {
		s.logger.Error("Error occurred in call to repository", zap.Error(err))
		return err
	}

	return nil
}

func (s service) DeleteGame(userID, gameID string) error {
	if err := s.repository.DeleteGame(userID, gameID); err != nil {
		s.logger.Error("Error occurred in call to repository", zap.Error(err))
		return err
	}

	return nil
}
