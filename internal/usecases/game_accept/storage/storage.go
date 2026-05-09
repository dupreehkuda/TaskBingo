package storage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

type Storage struct{ pool *pgxpool.Pool }

func New(pool *pgxpool.Pool) *Storage { return &Storage{pool: pool} }

// Accept marks the game started and stamps acceptance time. Limited to user2.
func (s *Storage) Accept(ctx context.Context, userID, gameID string) (pgconn.CommandTag, error) {
	return s.pool.Exec(ctx,
		`UPDATE games SET status = $1, accepted = $4 WHERE id = $2 AND user2_id = $3`,
		models.GameDBStarted, gameID, userID, time.Now(),
	)
}
