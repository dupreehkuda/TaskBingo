package storage

import (
	"context"
	"strings"
	"time"

	"go.uber.org/zap"
)

// CreateUser inserts new user's data in the database
func (s storage) CreateUser(userID, username, email, passwordHash, passwordSalt, city string) error {
	conn, err := s.pool.Acquire(context.Background())
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()

	ctx := context.Background()
	tx, err := conn.Begin(ctx)
	if err != nil {
		s.logger.Error("Error occurred creating tx", zap.Error(err))
		return err
	}

	tx.Exec(ctx, "INSERT INTO users (id, username, email, registered, city) VALUES ($1, $2, $3, $4, $5)", userID, strings.TrimSpace(username), strings.TrimSpace(email), time.Now(), strings.TrimSpace(city))
	tx.Exec(ctx, "INSERT INTO login (id, passwordhash, passwordsalt) VALUES ($1, $2, $3);", userID, passwordHash, passwordSalt)

	err = tx.Commit(ctx)
	if err != nil {
		s.logger.Error("Error occurred making reservation in db", zap.Error(err))
		return err
	}

	return nil
}
