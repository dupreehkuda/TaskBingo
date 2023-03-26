package storage

import (
	"context"
	"time"

	"go.uber.org/zap"
)

// RequestFriend requests friendship to user
func (s storage) RequestFriend(userID, friendID string) error {
	ctx := context.Background()
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)

	tx.Exec(ctx, "INSERT INTO friends (id, friend_id, status) VALUES ($1, $2, $3);", userID, friendID, Requested)
	tx.Exec(ctx, "INSERT INTO friends (id, friend_id, status) VALUES ($1, $2, $3);", friendID, userID, Response)

	err = tx.Commit(ctx)
	if err != nil {
		s.logger.Error("Error while committing tx", zap.Error(err))
		return err
	}

	return nil
}

// AcceptFriend accepts friendship request
func (s storage) AcceptFriend(userID, friendID string) error {
	ctx := context.Background()
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)

	tx.Exec(ctx, "UPDATE friends SET status = $1, since = $3 where id = $2 and friend_id = $4;", Friend, userID, time.Now(), friendID)
	tx.Exec(ctx, "UPDATE friends SET status = $1, since = $3 where id = $2 and friend_id = $4;", Friend, friendID, time.Now(), userID)

	err = tx.Commit(ctx)
	if err != nil {
		s.logger.Error("Error while committing tx", zap.Error(err))
		return err
	}

	return nil
}

// DeleteFriend deletes friendship or cancels request
func (s storage) DeleteFriend(userID, friendID string) error {
	ctx := context.Background()
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)

	tx.Exec(ctx, "DELETE FROM friends WHERE id = $1 and friend_id = $2;", userID, friendID)
	tx.Exec(ctx, "DELETE FROM friends WHERE id = $1 and friend_id = $2;", friendID, userID)

	err = tx.Commit(ctx)
	if err != nil {
		s.logger.Error("Error while committing tx", zap.Error(err))
		return err
	}

	return nil
}
