package storage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dupreehkuda/TaskBingo/pkg/pgxtx"
)

type Querier interface {
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

type Storage struct{ pool *pgxpool.Pool }

func New(pool *pgxpool.Pool) *Storage { return &Storage{pool: pool} }

func (s *Storage) q(ctx context.Context) Querier {
	if tx, ok := pgxtx.From(ctx); ok {
		return tx
	}
	return s.pool
}

// Insert writes the pack and appends its ID to the creator's likedpacks.
// Caller must wrap in pgxtx.RunInTx.
func (s *Storage) Insert(ctx context.Context, packID, creatorID, title string, tasks []string, isPrivate bool) error {
	if _, err := s.q(ctx).Exec(ctx,
		`INSERT INTO packs (id, title, tasks, creator, created, is_private) VALUES ($1, $2, $3, $4, $5, $6)`,
		packID, title, tasks, creatorID, time.Now(), isPrivate,
	); err != nil {
		return err
	}
	_, err := s.q(ctx).Exec(ctx,
		`UPDATE users SET likedpacks = ARRAY_APPEND(likedpacks, $1) WHERE id = $2`,
		packID, creatorID,
	)
	return err
}
