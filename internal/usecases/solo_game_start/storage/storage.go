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

// Create persists a solo game (kind='solo', user2_id NULL) and appends its ID
// to the player's games array. Caller must wrap in pgxtx.RunInTx.
func (s *Storage) Create(ctx context.Context, g *models.Game) error {
	if _, err := s.q(ctx).Exec(ctx,
		`INSERT INTO games
		    (id, user1_id, pack_id, status, numbers, user1_numbers, user2_numbers, created, kind)
		 VALUES ($1, $2, $3, $4, $5, $6, '{}'::integer[], $7, 'solo')`,
		g.GameID, g.User1Id, g.PackId, models.GameDBStarted,
		g.Numbers, g.User1Numbers, time.Now(),
	); err != nil {
		return err
	}
	_, err := s.q(ctx).Exec(ctx,
		`UPDATE users SET games = ARRAY_APPEND(games, $1) WHERE id = $2`,
		g.GameID, g.User1Id,
	)
	return err
}
