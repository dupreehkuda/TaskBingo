package processors

import (
	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
	"go.uber.org/zap"
)

// CreateGame creates new game instance in service
func (p processor) CreateGame(game *models.Game) error {
	if err := p.storage.CreateGame(game); err != nil {
		p.logger.Error("Error occurred in call to storage", zap.Error(err))
		return err
	}

	return nil
}
