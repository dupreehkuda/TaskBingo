package processors

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// SetTaskPack sets new task pack and assigns it to creator
func (p processor) SetTaskPack(login string, pack *models.TaskPack) error {
	err := p.taskStorage.SetTaskPack(pack)
	if err != nil {
		p.logger.Error("Error in call to task storage", zap.Error(err))
		return err
	}

	err = p.userStorage.AssignNewPack(login, pack.TaskID)
	if err != nil {
		p.logger.Error("Error in call to user storage", zap.Error(err))
		return err
	}

	return nil
}
