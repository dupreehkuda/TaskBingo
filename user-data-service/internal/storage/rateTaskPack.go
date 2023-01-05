package storage

import (
	"context"

	"go.uber.org/zap"
)

// RatePack rates pack by inc
func (s storage) RatePack(pack string, inc int) error {
	ctx := context.Background()
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()

	_, err = conn.Query(ctx, "UPDATE ratings SET rating = rating + $1 WHERE id = $2;", inc, pack)
	if err != nil {
		s.logger.Error("Error while executing query", zap.Error(err))
		return err
	}

	return nil
}
