package storage

import (
	"context"
	"time"

	"go.uber.org/zap"
)

// RequestFriend requests friendship to user
func (s storage) RequestFriend(login, person string) error {
	ctx := context.Background()
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)

	tx.Exec(ctx, "INSERT INTO friends (id, friend, status) VALUES ($1, $2, $3);", login, person, Requested)
	tx.Exec(ctx, "INSERT INTO friends (id, friend, status) VALUES ($1, $2, $3);", person, login, Responce)

	err = tx.Commit(ctx)
	if err != nil {
		s.logger.Error("Error while committing tx", zap.Error(err))
		return err
	}

	return nil
}

// AcceptFriend accepts friendship request
func (s storage) AcceptFriend(login, person string) error {
	ctx := context.Background()
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)

	tx.Exec(ctx, "UPDATE friends SET status = $1 and since = $4 where id = $2 and id = $3;", Friend, login, person, time.Now())

	err = tx.Commit(ctx)
	if err != nil {
		s.logger.Error("Error while committing tx", zap.Error(err))
		return err
	}

	return nil
}

// DeleteFriend deletes friendship or cancels request
func (s storage) DeleteFriend(login, person string) error {
	ctx := context.Background()
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)

	tx.Exec(ctx, "DELETE FROM friends WHERE id = $1 and id = $2;", login, person)

	err = tx.Commit(ctx)
	if err != nil {
		s.logger.Error("Error while committing tx", zap.Error(err))
		return err
	}

	return nil
}
