package storage

import (
	"context"
	"encoding/json"

	"go.uber.org/zap"
)

type response struct {
	UserID string `json:"userID"`
	Points int    `json:"points"`
}

func (s storage) Ping(userID string) ([]byte, error) {
	ctx := context.Background()
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return nil, err
	}
	defer conn.Release()

	var resp response

	err = conn.QueryRow(ctx, "SELECT nickname, wins FROM users WHERE id = $1", userID).Scan(&resp.UserID, &resp.Points)
	if err != nil {
		s.logger.Error("Error when executing statement", zap.Error(err))
		return nil, err
	}

	resultJSON, err := json.Marshal(resp)
	if err != nil {
		s.logger.Error("Error marshaling data", zap.Error(err))
		return nil, err
	}

	return resultJSON, nil
}
