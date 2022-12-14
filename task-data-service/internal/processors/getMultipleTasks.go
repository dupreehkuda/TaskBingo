package processors

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/models"
)

func (p processor) GetMultiplePacks(ids []string) (*[]models.TaskPack, error) {
	resp, err := p.storage.GetMultiplePacks(ids)
	if err != nil {
		p.logger.Error("Error occurred in call to storage", zap.Error(err))
		return nil, err
	}

	return resp, nil
}
