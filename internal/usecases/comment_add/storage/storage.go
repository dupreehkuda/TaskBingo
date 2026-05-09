package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

type Storage struct{ pool *pgxpool.Pool }

func New(pool *pgxpool.Pool) *Storage { return &Storage{pool: pool} }

// Insert writes a new comment row and returns the persisted record (with
// server-set timestamps).
func (s *Storage) Insert(ctx context.Context, c *models.Comment) (*models.Comment, error) {
	err := s.pool.QueryRow(ctx,
		`INSERT INTO game_comments (id, game_id, user_id, body)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id, game_id, user_id, body, created_at, COALESCE(updated_at, created_at)`,
		c.ID, c.GameID, c.UserID, c.Body,
	).Scan(&c.ID, &c.GameID, &c.UserID, &c.Body, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return c, nil
}
