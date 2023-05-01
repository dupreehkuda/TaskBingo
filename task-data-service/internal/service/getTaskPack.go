package service

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/models"
)

// GetTaskPack gets one task pack by taskId
func (s service) GetTaskPack(packID string) (*models.TaskPack, error) {
	resp, err := s.repository.GetTaskPack(packID)
	if err != nil {
		s.logger.Error("Error occurred in call to repository", zap.Error(err))
		return nil, err
	}

	return resp, nil
}
