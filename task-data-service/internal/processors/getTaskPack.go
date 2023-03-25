package processors

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/models"
)

// GetTaskPack gets one task pack by taskId
func (p processor) GetTaskPack(packID string) (*models.TaskPack, error) {
	resp, err := p.storage.GetTaskPack(packID)
	if err != nil {
		p.logger.Error("Error occurred in call to storage", zap.Error(err))
		return nil, err
	}

	return resp, nil
}
