package processors

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/models"
)

// AddTaskPack adds new task pack to storage
func (p processor) AddTaskPack(pack *models.TaskPack) error {
	err := p.storage.AddTaskPack(pack)
	if err != nil {
		p.logger.Error("Error when calling storage", zap.Error(err))
		return err
	}

	return nil
}
