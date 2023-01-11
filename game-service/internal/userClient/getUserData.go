package userClient

import (
	"context"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// GetUserData retrieves user data from user service
func (u userClient) GetUserData(login string) (*models.UserAccountInfoResponse, error) {
	resp, err := u.conn.GetUserData(context.Background(), &api.GetUserDataRequest{Login: login})
	if err != nil {
		u.logger.Error("Error when getting user data")
		return nil, err
	}

	res := models.UserAccountInfoResponse{
		Login:   resp.Login,
		City:    resp.City,
		Wins:    int(resp.Wins),
		Lose:    int(resp.Loses),
		Bingo:   int(resp.Bingo),
		Friends: []models.FriendsInfo{},
		Packs:   resp.Packs,
	}

	for _, val := range resp.Friends {
		res.Friends = append(res.Friends, models.FriendsInfo{
			Login: val.Login,
			City:  val.City,
			Bingo: val.Bingo,
		})
	}

	return &res, nil
}
