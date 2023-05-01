package service

import (
	"go.uber.org/zap"

	errs "github.com/dupreehkuda/TaskBingo/task-data-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/models"
)

// AddTaskPack adds new task pack to repository
func (s service) AddTaskPack(pack *models.TaskPack) error {
	exists, err := s.repository.CheckPackExistence(pack.ID)
	if err != nil {
		s.logger.Error("Error when calling repository", zap.Error(err))
		return err
	}

	if exists {
		return errs.ErrPackAlreadyExists
	}

	err = s.repository.AddTaskPack(pack)
	if err != nil {
		s.logger.Error("Error when calling repository", zap.Error(err))
		return err
	}

	return nil
}
