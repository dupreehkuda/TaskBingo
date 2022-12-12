package storage

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

func (s storage) GetUserData(login string) (models.GetUserDataResponse, error) {
	ctx := context.Background()
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return models.GetUserDataResponse{}, err
	}
	defer conn.Release()

	var resp models.GetUserDataResponse

	err = conn.QueryRow(ctx, "SELECT login, wins, email FROM users WHERE login = $1", login).Scan(&resp.UserID, &resp.Points, &resp.Email)
	if err != nil {
		s.logger.Error("Error when executing statement", zap.Error(err))
		return resp, err
	}

	return resp, nil
}
