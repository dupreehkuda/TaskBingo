package storage

import (
	"context"

	"go.uber.org/zap"
)

// LikePack likes or dislikes pack on user
func (s storage) LikePack(login, pack string, inc int) error {
	ctx := context.Background()
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)

	tx.Exec(ctx, "UPDATE ratings SET liked = liked + $1 WHERE id = $2;", inc, pack)

	if inc == 1 {
		// add pack id to the liked array
		tx.Exec(ctx, "UPDATE users SET packs = ARRAY_APPEND(packs, $1) WHERE login = $2;", pack, login)
	} else {
		// remove pack id from the liked array
		tx.Exec(ctx, "UPDATE users SET packs = ARRAY_REMOVE(packs, $1) WHERE login = $2;", pack, login)
	}

	err = tx.Commit(ctx)
	if err != nil {
		s.logger.Error("Error while committing tx", zap.Error(err))
		return err
	}

	return nil
}
