package repository

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// GetTaskPack retrieves a task pack from database
func (r repository) GetTaskPack(ctx context.Context, packId string) (*models.TaskPack, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		r.logger.Error("Error while acquiring connection", zap.Error(err))
		return nil, err
	}
	defer conn.Release()

	var resp models.TaskPack
	err = conn.QueryRow(ctx, "SELECT id, title, tasks FROM packs WHERE id = $1;", packId).
		Scan(&resp.ID, &resp.Pack.Title, &resp.Pack.Tasks)

	if err != nil {
		r.logger.Error("Error while executing query", zap.Error(err))
		return nil, err
	}

	return &resp, nil
}
