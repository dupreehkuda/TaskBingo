package repository

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// AddTaskPack adds task pack to the database
func (r repository) AddTaskPack(ctx context.Context, userID string, pack *models.TaskPack) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		r.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()

	createdAt := time.Now()

	tx, err := conn.Begin(ctx)

	_, err = tx.Exec(ctx, "INSERT INTO packs (id, title, tasks, creator, created) VALUES ($1, $2, $3, $4, $5);",
		pack.ID, pack.Pack.Title, pack.Pack.Tasks, userID, createdAt)

	_, err = tx.Exec(ctx, "INSERT INTO ratings (id, pack, creator, created) VALUES ($1, $2, $3, $4);",
		pack.ID, pack.Pack.Title, userID, createdAt)

	_, err = tx.Exec(ctx, "UPDATE users SET likedPacks = ARRAY_APPEND(likedPacks, $1) WHERE id = $2;",
		pack.ID, userID)

	if err != nil {
		r.logger.Error("Error query", zap.Error(err))
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		r.logger.Error("Error while committing tx", zap.Error(err))
		return err
	}

	return nil
}
