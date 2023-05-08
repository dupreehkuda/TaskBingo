package userRepository

import (
	"context"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// GetUserData retrieves user data from user service
func (u userRepository) GetUserData(userID string) (*models.UserAccountInfoResponse, error) {
	resp, err := u.conn.GetUserData(context.Background(), &api.GetUserDataRequest{UserID: userID})
	if err != nil {
		u.logger.Error("Error when getting user data")
		return nil, err
	}

	return mapFromUserDataResponse(resp), nil
}
