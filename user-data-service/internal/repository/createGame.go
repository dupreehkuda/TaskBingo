package repository

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// CreateGame writes new game to db and changes users game arrays
func (r repository) CreateGame(game *models.Game) error {
	ctx := context.Background()
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		r.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)

	tx.Exec(ctx, `INSERT INTO games (id, user1_id, user2_id, pack_id, status, created) VALUES ($1, $2, $3, $4, $5, $6)`,
		game.GameID, game.User1Id, game.User2Id, game.PackId, game.Status, time.Now())

	tx.Exec(ctx, `UPDATE users SET games = ARRAY_APPEND(games, $1) WHERE id = $2`, game.GameID, game.User1Id)
	tx.Exec(ctx, `UPDATE users SET games = ARRAY_APPEND(games, $1) WHERE id = $2`, game.GameID, game.User2Id)

	err = tx.Commit(ctx)
	if err != nil {
		r.logger.Error("Error while committing tx", zap.Error(err))
		return err
	}

	return nil
}
