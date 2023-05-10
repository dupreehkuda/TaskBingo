package service

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// todo should only go to repository

// SetTaskPack sets new task pack and assigns it to creator
func (s service) SetTaskPack(ctx context.Context, userID string, pack *models.TaskPack) error {
	packID, err := uuid.NewUUID()
	if err != nil {
		s.logger.Error("Unable to generate UUID", zap.Error(err))
		return err
	}

	pack.ID = packID.String()
	
	return s.repository.SetNewTaskPack(ctx, userID, pack)
}
