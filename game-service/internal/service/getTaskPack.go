package service

import (
	"context"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// todo should only go to repository

// GetTaskPack gets task pack by provided packID
func (s service) GetTaskPack(ctx context.Context, packID string) (*models.TaskPack, error) {
	return s.repository.GetTaskPack(ctx, packID)
}
