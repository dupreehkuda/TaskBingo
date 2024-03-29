package repository

import (
	"context"
	"time"

	"go.uber.org/zap"
)

// RequestFriend requests friendship to user
func (r repository) RequestFriend(ctx context.Context, userID, friendID string) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		r.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)

	tx.Exec(ctx, "INSERT INTO friends (id, friend_id, status) VALUES ($1, $2, $3);", userID, friendID, FriendRequested)
	tx.Exec(ctx, "INSERT INTO friends (id, friend_id, status) VALUES ($1, $2, $3);", friendID, userID, FriendResponse)

	err = tx.Commit(ctx)
	if err != nil {
		r.logger.Error("Error while committing tx", zap.Error(err))
		return err
	}

	return nil
}

// AcceptFriend accepts friendship request
func (r repository) AcceptFriend(ctx context.Context, userID, friendID string) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		r.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)

	tx.Exec(ctx, "UPDATE friends SET status = $1, since = $3 where id = $2 and friend_id = $4;", FriendFriend, userID, time.Now(), friendID)
	tx.Exec(ctx, "UPDATE friends SET status = $1, since = $3 where id = $2 and friend_id = $4;", FriendFriend, friendID, time.Now(), userID)

	err = tx.Commit(ctx)
	if err != nil {
		r.logger.Error("Error while committing tx", zap.Error(err))
		return err
	}

	return nil
}

// DeleteFriend deletes friendship or cancels request
func (r repository) DeleteFriend(ctx context.Context, userID, friendID string) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		r.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)

	tx.Exec(ctx, "DELETE FROM friends WHERE id = $1 and friend_id = $2;", userID, friendID)
	tx.Exec(ctx, "DELETE FROM friends WHERE id = $1 and friend_id = $2;", friendID, userID)

	err = tx.Commit(ctx)
	if err != nil {
		r.logger.Error("Error while committing tx", zap.Error(err))
		return err
	}

	return nil
}
