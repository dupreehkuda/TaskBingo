package repository

import (
	"context"

	"go.uber.org/zap"
)

// GetRatedPacks retrieves some rated packs from the db
func (r repository) GetRatedPacks(ctx context.Context) ([]string, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		r.logger.Error("Error while acquiring connection", zap.Error(err))
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, "SELECT id FROM ratings ORDER BY rating DESC;")
	if err != nil {
		r.logger.Error("Error while executing query", zap.Error(err))
		return nil, err
	}

	var resp []string
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			r.logger.Error("Error while scanning query", zap.Error(err))
			return nil, err
		}

		resp = append(resp, id)
	}

	return resp, nil
}
