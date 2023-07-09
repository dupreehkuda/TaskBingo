package service

import (
	"context"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// CreateGame creates new game instance in service
func (s service) CreateGame(ctx context.Context, game *models.Game) error {
	return s.repository.CreateGame(ctx, game)
}

// AcceptGame changes game status when user accepts it
func (s service) AcceptGame(ctx context.Context, userID, gameID string) error {
	return s.repository.AcceptGame(ctx, userID, gameID)
}

// DeleteGame deletes the game that user deleted/declined
func (s service) DeleteGame(ctx context.Context, userID, gameID string) error {
	return s.repository.DeleteGame(ctx, userID, gameID)
}
