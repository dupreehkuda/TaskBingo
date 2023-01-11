package storage

import (
	"context"

	"go.uber.org/zap"
)

// CheckDuplicateUser checks if user is already existing
func (s storage) CheckDuplicateUser(login, email string) (bool, error) {
	var (
		dbLogin string
		dbEmail string
	)

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

	conn.QueryRow(context.Background(), "SELECT email FROM users WHERE email=$1", email).Scan(&dbEmail)
	if dbEmail == email {
		return true, nil
	}

	return false, nil
}
