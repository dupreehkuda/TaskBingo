package service

import (
	"github.com/google/uuid"
	"go.uber.org/zap"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// SetTaskPack sets new task pack and assigns it to creator
func (s service) SetTaskPack(userID string, pack *models.TaskPack) error {
	packID, err := uuid.NewUUID()
	if err != nil {
		s.logger.Error("Unable to generate UUID", zap.Error(err))
		return err
	}

	pack.ID = packID.String()

	err = s.taskRepository.SetTaskPack(pack)
	if err != nil {
		if err == errs.ErrPackAlreadyExists {
			return err
		}

		s.logger.Error("Error in call to task repository", zap.Error(err))
		return err
	}

	err = s.userRepository.AssignNewPack(userID, packID.String(), pack.Pack.Title)
	if err != nil {
		s.logger.Error("Error in call to user repository", zap.Error(err))
		return err
	}

	return nil
}
