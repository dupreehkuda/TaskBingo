package repository

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// Enum for game status
const (
	_ = iota
	GameRequested
	GameStarted
	GameEnded
)

// CreateGame writes new game to db and changes users game arrays
func (r repository) CreateGame(ctx context.Context, game *models.Game) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		r.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)

	_, err = tx.Exec(ctx, `INSERT INTO games (id, user1_id, user2_id, pack_id, status, numbers, user1_numbers, user2_numbers, created) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		game.GameID, game.User1Id, game.User2Id, game.PackId, GameRequested, game.Numbers, game.User1Numbers, game.User2Numbers, time.Now())

	_, err = tx.Exec(ctx, `UPDATE users SET games = ARRAY_APPEND(games, $1) WHERE id = $2`, game.GameID, game.User1Id)
	_, err = tx.Exec(ctx, `UPDATE users SET games = ARRAY_APPEND(games, $1) WHERE id = $2`, game.GameID, game.User2Id)

	if err != nil {
		r.logger.Error("Error while executing statement", zap.Error(err))
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		r.logger.Error("Error while committing tx", zap.Error(err))
		return err
	}

	return nil
}

// AcceptGame updates a game when user accepts it
func (r repository) AcceptGame(ctx context.Context, userID, gameID string) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		r.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)

	_, err = tx.Exec(ctx, `UPDATE games SET status = $1, accepted = $4 WHERE id = $2 AND user2_id = $3;`, GameStarted, gameID, userID, time.Now())
	if err != nil {
		r.logger.Error("Error while executing statement", zap.Error(err))
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		r.logger.Error("Error while committing tx", zap.Error(err))
		return err
	}

	return nil
}

// DeleteGame deletes a game when user declines or deletes it
func (r repository) DeleteGame(ctx context.Context, userID, gameID string) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		r.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)

	_, err = tx.Exec(ctx, `DELETE FROM games WHERE id = $1 AND user1_id = $2 OR id = $1 AND user2_id = $2`, gameID, userID)
	if err != nil {
		r.logger.Error("Error while executing statement", zap.Error(err))
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		r.logger.Error("Error while committing tx", zap.Error(err))
		return err
	}

	return nil
}

// GetGame retrieves current game from db
func (r repository) GetGame(ctx context.Context, gameID string) (*models.Game, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		r.logger.Error("Error while acquiring connection", zap.Error(err))
		return nil, err
	}
	defer conn.Release()

	row := conn.QueryRow(ctx, "SELECT id, user1_id, user2_id, pack_id, status, user1_bingo, user2_bingo, numbers, user1_numbers, user2_numbers FROM games where id = $1;", gameID)
	if err != nil {
		r.logger.Error("Error when executing statement", zap.Error(err))
		return nil, err
	}

	resp := &models.Game{}
	err = row.Scan(&resp.GameID, &resp.User1Id, &resp.User2Id, &resp.PackId, &resp.Status, &resp.User1Bingo, &resp.User2Bingo, &resp.Numbers, &resp.User1Numbers, &resp.User2Numbers)
	if err != nil {
		r.logger.Error("Error scanning data", zap.Error(err))
		return nil, err
	}

	return resp, nil
}

// AchieveGame writes finished game data to db
func (r repository) AchieveGame(ctx context.Context, game *models.Game) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		r.logger.Error("Error while acquiring connection", zap.Error(err))
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)

	if game.Winner != "" {
		_, err = tx.Exec(ctx, `UPDATE games SET status = $1, user1_bingo = user1_bingo + $2, user2_bingo = user2_bingo + $3,
	                 winner = $4, user1_numbers = $5, user2_numbers = $6, finished = $7 WHERE id = $8;`,
			GameEnded, game.User1Bingo, game.User2Bingo, game.Winner, game.User1Numbers, game.User2Numbers, time.Now(), game.GameID)

		if err != nil {
			r.logger.Error("Error while executing statement", zap.Error(err))
			return err
		}
	} else {
		_, err = tx.Exec(ctx, `UPDATE games SET status = $1, user1_bingo = user1_bingo + $2, user2_bingo = user2_bingo + $3,
	                 user1_numbers = $4, user2_numbers = $5, finished = $6 WHERE id = $7;`,
			GameEnded, game.User1Bingo, game.User2Bingo, game.User1Numbers, game.User2Numbers, time.Now(), game.GameID)

		if err != nil {
			r.logger.Error("Error while executing statement", zap.Error(err))
			return err
		}
	}

	_, err = tx.Exec(ctx, `UPDATE users SET bingo = bingo + $1 WHERE id = $2;`, game.User1Bingo, game.User1Id)
	if err != nil {
		r.logger.Error("Error while executing statement", zap.Error(err))
		return err
	}

	_, err = tx.Exec(ctx, `UPDATE users SET bingo = bingo + $1 WHERE id = $2;`, game.User2Bingo, game.User2Id)
	if err != nil {
		r.logger.Error("Error while executing statement", zap.Error(err))
		return err
	}

	if game.User1Id == game.Winner {
		_, err = tx.Exec(ctx, `UPDATE users SET wins = wins + 1 WHERE id = $1;`, game.User1Id)
		_, err = tx.Exec(ctx, `UPDATE users SET lose = lose + 1 WHERE id = $1;`, game.User2Id)
	} else {
		_, err = tx.Exec(ctx, `UPDATE users SET wins = wins + 1 WHERE id = $1;`, game.User2Id)
		_, err = tx.Exec(ctx, `UPDATE users SET lose = lose + 1 WHERE id = $1;`, game.User1Id)
	}

	if err != nil {
		r.logger.Error("Error while executing statement", zap.Error(err))
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		r.logger.Error("Error while committing tx", zap.Error(err))
		return err
	}

	return nil
}
