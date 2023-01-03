package storage

import (
	"encoding/json"

	"go.uber.org/zap"

	errs "github.com/dupreehkuda/TaskBingo/task-data-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/models"
)

func (s storage) GetMultiplePacks(ids []string) (*[]models.TaskPack, error) {
	res, err := s.handle.JSONMGet(".", ids...)
	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, errs.ErrNoSuchPack
		}
		s.logger.Error("Error occurred when calling redis", zap.Error(err))
		return nil, err
	}

	var tasks []models.TaskPack
	packs := res.([]interface{})

	for i := range ids {
		var pack models.TaskPack
		err = json.Unmarshal(packs[i].([]byte), &pack)
		if err != nil {
			s.logger.Error("Error occurred when unmarshaling data", zap.Error(err))
			return nil, err
		}

		tasks = append(tasks, pack)
	}

	return &tasks, nil
}
