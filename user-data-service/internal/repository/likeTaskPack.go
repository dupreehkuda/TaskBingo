package repository

import (
	"context"

	"go.uber.org/zap"
)

// LikePack likes or dislikes pack on user
func (r repository) LikePack(ctx context.Context, userID, pack string, inc int) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		r.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)

	tx.Exec(ctx, "UPDATE packs SET liked = liked + $1 WHERE id = $2;", inc, pack)

	if inc == 1 {
		// add pack id to the liked array
		tx.Exec(ctx, "UPDATE users SET likedpacks = ARRAY_APPEND(likedpacks, $1) WHERE id = $2;", pack, userID)
	} else {
		// remove pack id from the liked array
		tx.Exec(ctx, "UPDATE users SET likedpacks = ARRAY_REMOVE(likedpacks, $1) WHERE id = $2;", pack, userID)
	}

	err = tx.Commit(ctx)
	if err != nil {
		r.logger.Error("Error while committing tx", zap.Error(err))
		return err
	}

	return nil
}
