package service

import (
	"context"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// GetAllUsers gets all users
func (s service) GetAllUsers(ctx context.Context) (*[]models.AllUsers, error) {
	return s.repository.GetAllUsers(ctx)
}
