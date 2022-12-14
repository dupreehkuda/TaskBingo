package storage

import (
	"context"

	"go.uber.org/zap"
)

// CheckDuplicateUser checks if user is already existing
func (s storage) CheckDuplicateUser(login string) (bool, error) {
	var dbLogin string

	conn, err := s.pool.Acquire(context.Background())
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return false, err
	}
	defer conn.Release()

	conn.QueryRow(context.Background(), "SELECT login FROM users WHERE login=$1", login).Scan(&dbLogin)

	if dbLogin == login {
		return true, nil
	}

	return false, nil
}
