package service

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

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

// GetTaskPack gets task pack by provided packID
func (s service) GetTaskPack(ctx context.Context, packID string) (*models.TaskPack, error) {
	return s.repository.GetTaskPack(ctx, packID)
}

// GetRatedPacks gets some most rated packs
func (s service) GetRatedPacks(ctx context.Context) (*models.Packs, error) {
	return s.repository.GetRatedPacks(ctx)
}

// LikeTaskPack likes pack by user
func (s service) LikeTaskPack(ctx context.Context, userID, pack string) error {
	return s.repository.LikeTaskPack(ctx, userID, pack)
}

// DislikeTaskPack dislikes pack by user
func (s service) DislikeTaskPack(ctx context.Context, userID, pack string) error {
	return s.repository.DislikeTaskPack(ctx, userID, pack)
}

// RateTaskPack adds to pack rating
func (s service) RateTaskPack(ctx context.Context, userID, pack string) error {
	return s.repository.RateTaskPack(ctx, userID, pack)
}

// UnrateTaskPack removes from pack rating
func (s service) UnrateTaskPack(ctx context.Context, userID, pack string) error {
	return s.repository.UnrateTaskPack(ctx, userID, pack)
}
