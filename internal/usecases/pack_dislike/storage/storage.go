package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

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

// Dislike decrements packs.liked and removes the pack from users.likedpacks.
// Caller must wrap in pgxtx.RunInTx.
func (s *Storage) Dislike(ctx context.Context, userID, packID string) error {
	if _, err := s.q(ctx).Exec(ctx,
		`UPDATE packs SET liked = liked - 1 WHERE id = $1`, packID,
	); err != nil {
		return err
	}
	_, err := s.q(ctx).Exec(ctx,
		`UPDATE users SET likedpacks = ARRAY_REMOVE(likedpacks, $1) WHERE id = $2`,
		packID, userID,
	)
	return err
}
