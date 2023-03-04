package storage

import (
	"context"

	"go.uber.org/zap"
)

// CheckDuplicateUser checks if user is already existing
func (s storage) CheckDuplicateUser(username, email string) (bool, error) {
	var (
		dbUsername string
		dbEmail    string
	)

	conn, err := s.pool.Acquire(context.Background())
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return false, err
	}
	defer conn.Release()

	conn.QueryRow(context.Background(), "SELECT username FROM users WHERE username = $1", username).Scan(&dbUsername)
	if dbUsername == username {
		s.logger.Debug("wtf? 1", zap.String("db", dbUsername), zap.String("req", username))
		return true, nil
	}

	conn.QueryRow(context.Background(), "SELECT email FROM users WHERE email = $1", email).Scan(&dbEmail)
	if dbEmail == email {
		return true, nil
	}

	return false, nil
}
