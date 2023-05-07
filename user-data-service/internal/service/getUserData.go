package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// GetUserData gets some important account data
func (s service) GetUserData(ctx context.Context, userId string) (*models.GetUserDataResponse, error) {
	resp, err := s.repository.GetUserData(ctx, userId)
	if err != nil {
		s.logger.Error("Unable to call repository", zap.Error(err))
		return nil, err
	}

	return resp, nil
}
