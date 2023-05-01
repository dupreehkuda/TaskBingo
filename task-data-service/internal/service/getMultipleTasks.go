package service

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/models"
)

func (s service) GetMultiplePacks(ids []string) (*[]models.TaskPack, error) {
	resp, err := s.repository.GetMultiplePacks(ids)
	if err != nil {
		s.logger.Error("Error occurred in call to repository", zap.Error(err))
		return nil, err
	}

	return resp, nil
}
