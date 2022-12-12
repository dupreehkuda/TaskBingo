package storage

import (
	"context"
	"time"

	"go.uber.org/zap"
)

// CreateUser inserts new user's data in the database
func (s storage) CreateUser(login, email, passwordHash, passwordSalt string) error {
	conn, err := s.pool.Acquire(context.Background())
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()

	// todo: id should not be the same with login, think about google/uuid

	ctx := context.Background()
	tx, err := conn.Begin(ctx)
	if err != nil {
		s.logger.Error("Error occurred creating tx", zap.Error(err))
		return err
	}

	tx.Exec(ctx, "INSERT INTO users (id, login, email, registered) VALUES ($1, $1, $2, $3)", login, email, time.Now())
	tx.Exec(ctx, "INSERT INTO login (id, passwordhash, passwordsalt) VALUES ($1, $2, $3);", login, passwordHash, passwordSalt)

	err = tx.Commit(ctx)
	if err != nil {
		s.logger.Error("Error occurred making reservation in db", zap.Error(err))
		return err
	}

	return nil
}
