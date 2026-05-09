package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct{ pool *pgxpool.Pool }

func New(pool *pgxpool.Pool) *Storage { return &Storage{pool: pool} }

// Delete removes the comment if owned by userID. Caller checks RowsAffected
// to map zero rows to ErrNotFound.
func (s *Storage) Delete(ctx context.Context, commentID, userID string) (pgconn.CommandTag, error) {
	return s.pool.Exec(ctx,
		`DELETE FROM game_comments WHERE id = $1 AND user_id = $2`,
		commentID, userID,
	)
}
