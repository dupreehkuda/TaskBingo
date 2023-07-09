package service

import (
	"context"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// AddTaskPack sets and assigns new task pack
func (s service) AddTaskPack(ctx context.Context, userID string, pack *models.TaskPack) error {
	return s.repository.AddTaskPack(ctx, userID, pack)
}

// GetTaskPack gets requested pack
func (s service) GetTaskPack(ctx context.Context, packId string) (*models.TaskPack, error) {
	return s.repository.GetTaskPack(ctx, packId)
}

// GetRatedPacks gets some of most rated packs
func (s service) GetRatedPacks(ctx context.Context) (*[]models.TaskPack, error) {
	return s.repository.GetRatedPacks(ctx)
}

// LikePack likes or dislikes the pack by inc
func (s service) LikePack(ctx context.Context, userID, pack string, inc int) error {
	return s.repository.LikePack(ctx, userID, pack, inc)
}

// RatePack rates pack by inc
func (s service) RatePack(ctx context.Context, userID, pack string, inc int) error {
	return s.repository.RatePack(ctx, userID, pack, inc)
}
