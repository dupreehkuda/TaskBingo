package service

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// GetTaskPack gets task pack by provided packID
func (s service) GetTaskPack(packID string) (*models.TaskPack, error) {
	resp, err := s.taskRepository.GetTaskPack(packID)
	if err != nil {
		s.logger.Error("Error occurred in call to task repository", zap.Error(err))
		return nil, err
	}

	return resp, nil
}
