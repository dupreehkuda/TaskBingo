package processors

import (
	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
	"go.uber.org/zap"
)

func (p processor) GetUserData(userId string) (*models.GetUserDataResponse, error) {
	resp, err := p.storage.GetUserData(userId)
	if err != nil {
		p.logger.Error("Unable to call storage", zap.Error(err))
		return nil, err
	}

	return &resp, nil
}
