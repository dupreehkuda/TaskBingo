package repository

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/models"
)

// AddTaskPack adds task pack to the database
func (r repository) AddTaskPack(pack *models.TaskPack) error {
	_, err := r.handle.JSONSet(pack.ID, ".", pack.Pack)
	if err != nil {
		r.logger.Error("Error in call to redis", zap.Error(err))
		return err
	}

	return nil
}
