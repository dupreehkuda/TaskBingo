package storage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/pkg/pgxtx"
)

type Querier interface {
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

type Storage struct{ pool *pgxpool.Pool }

func New(pool *pgxpool.Pool) *Storage { return &Storage{pool: pool} }

func (s *Storage) q(ctx context.Context) Querier {
	if tx, ok := pgxtx.From(ctx); ok {
		return tx
	}
	return s.pool
}

// Create inserts the game and appends its ID to both users' game arrays.
// Caller must wrap in pgxtx.RunInTx.
//
// Note: status is persisted as GameDBStarted to mirror legacy behaviour
// (the legacy CreateGame had a TODO that should have been GameDBRequested
// but accept-flow was disabled, so new games go straight to "started").
func (s *Storage) Create(ctx context.Context, game *models.Game) error {
	if _, err := s.q(ctx).Exec(ctx,
		`INSERT INTO games
            (id, user1_id, user2_id, pack_id, status, numbers, user1_numbers, user2_numbers, created)
         VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		game.GameID, game.User1Id, game.User2Id, game.PackId,
		models.GameDBStarted, game.Numbers, game.User1Numbers, game.User2Numbers, time.Now(),
	); err != nil {
		return err
	}
	if _, err := s.q(ctx).Exec(ctx,
		`UPDATE users SET games = ARRAY_APPEND(games, $1) WHERE id = $2`,
		game.GameID, game.User1Id,
	); err != nil {
		return err
	}
	_, err := s.q(ctx).Exec(ctx,
		`UPDATE users SET games = ARRAY_APPEND(games, $1) WHERE id = $2`,
		game.GameID, game.User2Id,
	)
	return err
}
