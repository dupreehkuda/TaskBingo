package repository

import (
	"context"

	"go.uber.org/zap"
)

// RatePack rates pack by inc
func (r repository) RatePack(ctx context.Context, userID, pack string, inc int) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		r.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)

	tx.Exec(ctx, "UPDATE ratings SET rating = rating + $1 WHERE id = $2;", inc, pack)

	if inc == 1 {
		tx.Exec(ctx, "UPDATE users SET ratedpacks = ARRAY_APPEND(ratedpacks, $1) WHERE id = $2;", pack, userID)
	} else {
		tx.Exec(ctx, "UPDATE users SET ratedpacks = ARRAY_REMOVE(ratedpacks, $1) WHERE id = $2;", pack, userID)
	}

	err = tx.Commit(ctx)
	if err != nil {
		r.logger.Error("Error while committing tx", zap.Error(err))
		return err
	}

	return nil
}
