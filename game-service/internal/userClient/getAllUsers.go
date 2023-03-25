package userClient

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// GetAllUsers calls user service for all users
func (u userClient) GetAllUsers() (*[]models.User, error) {
	resp, err := u.conn.GetAllPeople(context.Background(), &api.Empty{})
	if err != nil {
		u.logger.Error("Error when getting user data", zap.Error(err))
		return nil, err
	}

	var users []models.User

	for _, person := range resp.Person {
		users = append(users, models.User{
			UserID:   person.UserID,
			Username: person.Username,
			City:     person.City,
			Bingo:    int(person.Bingo),
		})
	}

	return &users, nil
}
