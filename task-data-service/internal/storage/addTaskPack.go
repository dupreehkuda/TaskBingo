package storage

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/models"
)

func (s storage) AddTaskPack(pack *models.TaskPack) error {
	_, err := s.handle.JSONSet(pack.TaskID, ".", pack.Tasks)
	if err != nil {
		s.logger.Error("Error in call to redis", zap.Error(err))
		return err
	}

	return nil
}
