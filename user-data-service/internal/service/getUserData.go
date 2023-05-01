package service

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// GetUserData gets some important account data
func (s service) GetUserData(userId string) (*models.GetUserDataResponse, error) {
	resp, err := s.repository.GetUserData(userId)
	if err != nil {
		s.logger.Error("Unable to call repository", zap.Error(err))
		return nil, err
	}

	return resp, nil
}
