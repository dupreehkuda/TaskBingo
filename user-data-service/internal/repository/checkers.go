package repository

import (
	"context"

	"go.uber.org/zap"
)

// CheckDuplicateUser checks if user is already existing
func (r repository) CheckDuplicateUser(ctx context.Context, username, email string) (bool, error) {
	var (
		dbUsername string
		dbEmail    string
	)

	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		r.logger.Error("Error while acquiring connection", zap.Error(err))
		return false, err
	}
	defer conn.Release()

	conn.QueryRow(ctx, "SELECT username FROM users WHERE username = $1", username).Scan(&dbUsername)
	if dbUsername == username {
		return true, nil
	}

	conn.QueryRow(ctx, "SELECT email FROM users WHERE email = $1", email).Scan(&dbEmail)
	if dbEmail == email {
		return true, nil
	}

	return false, nil
}
