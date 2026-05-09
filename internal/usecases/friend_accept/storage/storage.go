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

// Accept flips both friend rows to friend status. Caller must wrap in pgxtx.RunInTx.
func (s *Storage) Accept(ctx context.Context, userID, friendID string) error {
	now := time.Now()
	if _, err := s.q(ctx).Exec(ctx,
		`UPDATE friends SET status = $1, since = $2 WHERE id = $3 AND friend_id = $4`,
		models.FriendFriend, now, userID, friendID,
	); err != nil {
		return err
	}
	_, err := s.q(ctx).Exec(ctx,
		`UPDATE friends SET status = $1, since = $2 WHERE id = $3 AND friend_id = $4`,
		models.FriendFriend, now, friendID, userID,
	)
	return err
}
