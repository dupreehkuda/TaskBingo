package processors

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

func (p processor) GetTaskPack(packID string) (*models.TaskPack, error) {
	resp, err := p.taskStorage.GetTaskPack(packID)
	if err != nil {
		p.logger.Error("Error occurred in call to task storage", zap.Error(err))
		return nil, err
	}

	return resp, nil
}
