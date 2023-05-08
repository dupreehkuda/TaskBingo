package userRepository

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// GetAllUsers calls user service for all users
func (u userRepository) GetAllUsers(ctx context.Context) (*[]models.User, error) {
	resp, err := u.conn.GetAllPeople(ctx, &api.Empty{})
	if err != nil {
		u.logger.Error("Error when getting user data", zap.Error(err))
		return nil, err
	}

	return mapFromPeople(resp), nil
}
