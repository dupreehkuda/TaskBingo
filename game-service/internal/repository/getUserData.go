package repository

import (
	"context"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// GetUserData retrieves user data from user service
func (r repository) GetUserData(ctx context.Context, userID string) (*models.UserAccountInfoResponse, error) {
	resp, err := r.conn.GetUserData(ctx, &api.GetUserDataRequest{UserID: userID})
	if err != nil {
		r.logger.Error("Error when getting user data")
		return nil, err
	}

	return mapFromUserDataResponse(resp), nil
}
