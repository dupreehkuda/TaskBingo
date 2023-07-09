package service

import (
	"context"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// GetUserData gets user's most important info
func (s service) GetUserData(ctx context.Context, userID string) (*models.UserAccountInfo, error) {
	return s.repository.GetUserData(ctx, userID)
}
