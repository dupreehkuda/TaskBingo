package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

type Storage struct{ pool *pgxpool.Pool }

func New(pool *pgxpool.Pool) *Storage { return &Storage{pool: pool} }

// List returns the requester's own comments on the given game, oldest-first.
func (s *Storage) List(ctx context.Context, gameID, userID string) ([]models.Comment, error) {
	rows, err := s.pool.Query(ctx,
		`SELECT id, game_id, user_id, body, created_at, COALESCE(updated_at, created_at)
		 FROM game_comments
		 WHERE game_id = $1 AND user_id = $2
		 ORDER BY created_at`,
		gameID, userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []models.Comment
	for rows.Next() {
		var c models.Comment
		if err := rows.Scan(&c.ID, &c.GameID, &c.UserID, &c.Body, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, rows.Err()
}
