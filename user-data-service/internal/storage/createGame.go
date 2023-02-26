package storage

import (
	"context"
	"time"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
	"go.uber.org/zap"
)

// CreateGame writes new game to db and changes users game arrays
func (s storage) CreateGame(game *models.Game) error {
	ctx := context.Background()
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)

	tx.Exec(ctx, `INSERT INTO games (uuid, user1_id, user2_id, pack_id, status, created) VALUES ($1, $2, $3, $4, $5, $6)`,
		game.GameID, game.User1Id, game.User2Id, game.PackId, game.Status, time.Now())

	tx.Exec(ctx, `UPDATE users SET createdgames = ARRAY_APPEND(createdgames, $1) WHERE login = $2`, game.GameID, game.User1Id)
	tx.Exec(ctx, `UPDATE users SET requestedgames = ARRAY_APPEND(requestedgames, $1) WHERE login = $2`, game.GameID, game.User2Id)

	err = tx.Commit(ctx)
	if err != nil {
		s.logger.Error("Error while committing tx", zap.Error(err))
		return err
	}

	return nil
}
