package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// GetAllUsers gets all users from user service
func (s service) GetAllUsers(ctx context.Context) (*models.Users, error) {
	resp, err := s.repository.GetAllUsers(ctx)
	if err != nil {
		s.logger.Error("Error when calling user service", zap.Error(err))
		return nil, err
	}

	return resp, nil
}
