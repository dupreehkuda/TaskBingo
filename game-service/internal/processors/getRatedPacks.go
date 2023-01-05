package processors

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// GetRatedPacks gets some most rated packs
func (p processor) GetRatedPacks() (*[]models.TaskPack, error) {
	rated, err := p.userStorage.GetRatedPacks()
	if err != nil {
		p.logger.Error("Error occurred in call to user storage", zap.Error(err))
		return nil, err
	}

	packs, err := p.taskStorage.GetMultiplePacks(rated)
	if err != nil {
		p.logger.Error("Error occurred in call to task storage", zap.Error(err))
		return nil, err
	}

	return packs, nil
}
