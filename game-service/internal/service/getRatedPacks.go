package service

import (
	"context"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// todo should only go to repository

// GetRatedPacks gets some most rated packs
func (s service) GetRatedPacks(ctx context.Context) (*models.Packs, error) {
	return s.repository.GetRatedPacks(ctx)
}
