package repository

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// GetTaskPacks retrieves a task pack from database
func (r repository) GetTaskPacks(ctx context.Context, packIDs ...string) (*[]models.TaskPack, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		r.logger.Error("Error while acquiring connection", zap.Error(err))
		return nil, err
	}
	defer conn.Release()

	var resp []models.TaskPack

	rows, err := conn.Query(ctx, "SELECT id, title, tasks FROM packs WHERE id = ANY ($1);", packIDs)
	if err != nil {
		r.logger.Error("Error while executing query", zap.Error(err))
		return nil, err
	}

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
