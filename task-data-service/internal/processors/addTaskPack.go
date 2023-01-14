package processors

import (
	"go.uber.org/zap"

	errs "github.com/dupreehkuda/TaskBingo/task-data-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/models"
)

// AddTaskPack adds new task pack to storage
func (p processor) AddTaskPack(pack *models.TaskPack) error {
	exists, err := p.storage.CheckPackExistence(pack.TaskID)
	if err != nil {
		p.logger.Error("Error when calling storage", zap.Error(err))
		return err
	}

	if exists {
		return errs.ErrPackAlreadyExists
	}

	err = p.storage.AddTaskPack(pack)
	if err != nil {
		p.logger.Error("Error when calling storage", zap.Error(err))
		return err
	}

	return nil
}
