package storage

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// GetUserData retrieves user data from database
func (s storage) GetUserData(login string) (models.GetUserDataResponse, error) {
	ctx := context.Background()
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return models.GetUserDataResponse{}, err
	}
	defer conn.Release()

	var resp models.GetUserDataResponse

	row := conn.QueryRow(ctx, "SELECT login, city, wins, lose, scoreboard, friends, packs FROM users WHERE login = $1", login)
	err = row.Scan(&resp.Login, &resp.City, &resp.Wins, &resp.Lose, &resp.Scoreboard, &resp.Friends, &resp.Packs)
	if err != nil {
		s.logger.Error("Error when executing statement", zap.Error(err))
		return resp, err
	}

	return resp, nil
}
