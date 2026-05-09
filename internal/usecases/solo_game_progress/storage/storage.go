package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct{ pool *pgxpool.Pool }

func New(pool *pgxpool.Pool) *Storage { return &Storage{pool: pool} }

// Update persists the player's number set on a solo game. Ownership is enforced
// in the WHERE clause; non-owner or non-solo updates affect zero rows.
func (s *Storage) Update(ctx context.Context, userID, gameID string, userNumbers []int32) (pgconn.CommandTag, error) {
	return s.pool.Exec(ctx,
		`UPDATE games SET user1_numbers = $1
		 WHERE id = $2 AND user1_id = $3 AND kind = 'solo'`,
		userNumbers, gameID, userID,
	)
}
