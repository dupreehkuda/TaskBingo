package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// GetAllUsers gets all users
func (s service) GetAllUsers(ctx context.Context) (*[]models.AllUsers, error) {
	resp, err := s.repository.GetAllUsers(ctx)
	if err != nil {
		s.logger.Error("Error occurred in call to repository", zap.Error(err))
		return nil, err
	}

	return resp, nil
}
