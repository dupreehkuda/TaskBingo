package processors

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

func (p processor) GetUserData(login string) (*models.Response, error) {
	resp, err := p.userStorage.GetUserData(login)
	if err != nil {
		p.logger.Error("Error occurred in call to user service", zap.Error(err))
		return nil, err
	}

	return resp, nil
}
