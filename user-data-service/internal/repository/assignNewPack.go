package repository

import (
	"context"
	"time"

	"go.uber.org/zap"
)

// AssignNewPack assigns fresh pack to creator and adds pack to ratings
func (r repository) AssignNewPack(userID, packID, packName string) error {
	ctx := context.Background()
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		r.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)

	tx.Exec(ctx, "INSERT INTO ratings (id, pack, creator, created) VALUES ($1, $2, $3, $4);", packID, packName, userID, time.Now())
	tx.Exec(ctx, "UPDATE users SET likedPacks = ARRAY_APPEND(likedPacks, $1) WHERE id = $2;", packID, userID)

	err = tx.Commit(ctx)
	if err != nil {
		r.logger.Error("Error while committing tx", zap.Error(err))
		return err
	}

	return nil
}
