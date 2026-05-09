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

// Delete removes both mirrored rows. Caller must wrap in pgxtx.RunInTx.
func (s *Storage) Delete(ctx context.Context, userID, friendID string) error {
	if _, err := s.q(ctx).Exec(ctx,
		`DELETE FROM friends WHERE id = $1 AND friend_id = $2`, userID, friendID,
	); err != nil {
		return err
	}
	_, err := s.q(ctx).Exec(ctx,
		`DELETE FROM friends WHERE id = $1 AND friend_id = $2`, friendID, userID,
	)
	return err
}
