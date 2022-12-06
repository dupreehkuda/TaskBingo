package storage

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

func (s storage) Ping(userID string) (models.Response, error) {
	ctx := context.Background()
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return models.Response{}, err
	}
	defer conn.Release()

	var resp models.Response

	err = conn.QueryRow(ctx, "SELECT nickname, wins, email FROM users WHERE id = $1", userID).Scan(&resp.UserID, &resp.Points, &resp.Email)
	if err != nil {
		s.logger.Error("Error when executing statement", zap.Error(err))
		return resp, err
	}

	return resp, nil
}
