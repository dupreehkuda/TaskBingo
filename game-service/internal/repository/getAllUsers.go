package repository

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// GetAllUsers calls user service for all users
func (r repository) GetAllUsers(ctx context.Context) (*models.Users, error) {
	resp, err := r.conn.GetAllPeople(ctx, &api.Empty{})
	if err != nil {
		r.logger.Error("Error when getting user data", zap.Error(err))
		return nil, err
	}

	return mapFromPeople(resp), nil
}
