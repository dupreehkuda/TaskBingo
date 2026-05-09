package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct{ pool *pgxpool.Pool }

func New(pool *pgxpool.Pool) *Storage { return &Storage{pool: pool} }

// Delete removes the game row, but only if userID is one of its players.
// Mirrors the legacy ownership predicate.
func (s *Storage) Delete(ctx context.Context, userID, gameID string) (pgconn.CommandTag, error) {
	return s.pool.Exec(ctx,
		`DELETE FROM games WHERE id = $1 AND (user1_id = $2 OR user2_id = $2)`,
		gameID, userID,
	)
}
