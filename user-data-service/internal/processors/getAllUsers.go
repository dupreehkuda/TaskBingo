package processors

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// GetAllUsers gets all users
func (p processor) GetAllUsers() (*[]models.AllUsers, error) {
	resp, err := p.storage.GetAllUsers()
	if err != nil {
		p.logger.Error("Error occurred in call to storage", zap.Error(err))
		return nil, err
	}

	return resp, nil
}