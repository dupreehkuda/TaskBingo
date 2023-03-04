package storage

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// RatePack rates pack by inc
func (s storage) RatePack(userID, pack uuid.UUID, inc int) error {
	ctx := context.Background()
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
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
		s.logger.Error("Error while committing tx", zap.Error(err))
		return err
	}

	return nil
}
