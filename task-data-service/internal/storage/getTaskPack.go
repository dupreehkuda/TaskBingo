package storage

import (
	"encoding/json"

	"go.uber.org/zap"

	errs "github.com/dupreehkuda/TaskBingo/task-data-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/models"
)

func (s storage) GetTaskPack(taskId string) (*models.TaskPack, error) {
	res, err := s.handle.JSONGet(taskId, ".")
	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, errs.ErrNoSuchPack
		}
		s.logger.Error("Error occurred when calling redis", zap.Error(err))
		return nil, err
	}

	var tasks models.TaskPack
	err = json.Unmarshal(res.([]byte), &tasks.Tasks)
	if err != nil {
		s.logger.Error("Error occurred when unmarshaling data", zap.Error(err))
		return nil, err
	}

	tasks.TaskID = taskId

	return &tasks, nil
}
