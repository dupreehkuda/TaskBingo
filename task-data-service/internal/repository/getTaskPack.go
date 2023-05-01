package repository

import (
	"encoding/json"

	"go.uber.org/zap"

	errs "github.com/dupreehkuda/TaskBingo/task-data-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/models"
)

// GetTaskPack retrieves a task pack from database
func (r repository) GetTaskPack(taskId string) (*models.TaskPack, error) {
	res, err := r.handle.JSONGet(taskId, ".")
	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, errs.ErrNoSuchPack
		}
		r.logger.Error("Error occurred when calling redis", zap.Error(err))
		return nil, err
	}

	var tasks models.TaskPack
	err = json.Unmarshal(res.([]byte), &tasks.Pack)
	if err != nil {
		r.logger.Error("Error occurred when unmarshaling data", zap.Error(err))
		return nil, err
	}

	tasks.ID = taskId
	return &tasks, nil
}
