package service

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// GetAllUsers gets all users from user service
func (s service) GetAllUsers() (*[]models.User, error) {
	resp, err := s.userRepository.GetAllUsers()
	if err != nil {
		s.logger.Error("Error when calling user service", zap.Error(err))
		return nil, err
	}

	return resp, nil
}
