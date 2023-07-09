package repository

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// GetRatedPacks retrieves some rated packs from the db
func (r repository) GetRatedPacks(ctx context.Context) (*[]models.TaskPack, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		r.logger.Error("Error while acquiring connection", zap.Error(err))
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, "SELECT id, title, tasks FROM packs ORDER BY rating DESC;")
	if err != nil {
		r.logger.Error("Error while executing query", zap.Error(err))
		return nil, err
	}

	var resp []models.TaskPack

	for rows.Next() {
		var pack models.TaskPack

		err = rows.Scan(&pack.ID, &pack.Pack.Title, &pack.Pack.Tasks)
		if err != nil {
			r.logger.Error("Error while scanning query", zap.Error(err))
			return nil, err
		}

		resp = append(resp, pack)
	}

	return &resp, nil
}
