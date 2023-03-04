package storage

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// GetRatedPacks retrieves some rated packs from the db
func (s storage) GetRatedPacks() ([]uuid.UUID, error) {
	ctx := context.Background()
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, "SELECT id FROM ratings ORDER BY rating DESC;")
	if err != nil {
		s.logger.Error("Error while executing query", zap.Error(err))
		return nil, err
	}

	var resp []string

	for rows.Next() {
		var r string
		err = rows.Scan(&r)
		if err != nil {
			s.logger.Error("Error while scanning query", zap.Error(err))
			return nil, err
		}

		resp = append(resp, r)
	}

	// TODO Transforming strings to uuid here, uuid to strings in handlers
	var ans []uuid.UUID
	for _, val := range resp {
		id, err := uuid.Parse(val)
		if err != nil {
			s.logger.Error("Unable to parse UUID from db", zap.Error(err))
			return nil, err
		}

		ans = append(ans, id)
	}

	return ans, nil
}
