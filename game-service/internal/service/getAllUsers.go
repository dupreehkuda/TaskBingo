package service

import (
	"context"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// GetAllUsers gets all users from user service
func (s service) GetAllUsers(ctx context.Context) (*models.Users, error) {
	return s.repository.GetAllUsers(ctx)
}
