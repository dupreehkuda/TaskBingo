package userClient

import (
	"context"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// GetUserData retrieves user data from user service
func (u userClient) GetUserData(login string) (*models.Response, error) {
	resp, err := u.conn.GetUserData(context.Background(), &api.GetUserDataRequest{Login: login})
	if err != nil {
		u.logger.Error("Error when getting user data")
		return nil, err
	}

	return &models.Response{
		UserID: resp.Login,
		Points: int(resp.Points),
		Email:  resp.Email,
	}, nil
}
