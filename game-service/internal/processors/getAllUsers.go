package processors

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// GetAllUsers gets all users from user service
func (p processor) GetAllUsers() (*[]models.User, error) {
	resp, err := p.userStorage.GetAllUsers()
	if err != nil {
		p.logger.Error("Error when calling user service", zap.Error(err))
		return nil, err
	}

	return resp, nil
}
