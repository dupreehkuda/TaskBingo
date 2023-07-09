package service

import (
	"context"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// GetGame gets current game from repository
func (s service) GetGame(ctx context.Context, gameID string) (*models.Game, error) {
	return s.repository.GetGame(ctx, gameID)
}

// AchieveGame writes game data when its finished
func (s service) AchieveGame(ctx context.Context, game *models.Game) error {
	return s.repository.AchieveGame(ctx, game)
}
