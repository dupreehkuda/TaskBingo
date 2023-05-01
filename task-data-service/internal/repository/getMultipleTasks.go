package repository

import (
	"encoding/json"

	"go.uber.org/zap"

	errs "github.com/dupreehkuda/TaskBingo/task-data-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/models"
)

func (r repository) GetMultiplePacks(ids []string) (*[]models.TaskPack, error) {
	res, err := r.handle.JSONMGet(".", ids...)
	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, errs.ErrNoSuchPack
		}
		r.logger.Error("Error occurred when calling redis", zap.Error(err))
		return nil, err
	}

	var tasks []models.TaskPack
	packs := res.([]interface{})

	for i := range ids {
		var pack models.TaskPack
		err = json.Unmarshal(packs[i].([]byte), &pack.Pack)
		if err != nil {
			r.logger.Error("Error occurred when unmarshaling data", zap.Error(err))
			return nil, err
		}

		pack.ID = ids[i]
		tasks = append(tasks, pack)
	}

	return &tasks, nil
}
