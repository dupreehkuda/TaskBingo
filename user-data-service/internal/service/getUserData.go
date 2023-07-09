package service

import (
	"context"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// GetUserData gets some important account data
func (s service) GetUserData(ctx context.Context, userId string) (*models.GetUserDataResponse, error) {
	return s.repository.GetUserData(ctx, userId)
}
