package processors

import (
	"github.com/google/uuid"
	"go.uber.org/zap"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// SetTaskPack sets new task pack and assigns it to creator
func (p processor) SetTaskPack(userID string, pack *models.TaskPack) error {
	packID, err := uuid.NewUUID()
	if err != nil {
		p.logger.Error("Unable to generate UUID", zap.Error(err))
		return err
	}

	pack.ID = packID.String()

	err = p.taskStorage.SetTaskPack(pack)
	if err != nil {
		if err == errs.ErrPackAlreadyExists {
			return err
		}

		p.logger.Error("Error in call to task storage", zap.Error(err))
		return err
	}

	err = p.userStorage.AssignNewPack(userID, packID.String(), pack.Pack.Title)
	if err != nil {
		p.logger.Error("Error in call to user storage", zap.Error(err))
		return err
	}

	return nil
}
