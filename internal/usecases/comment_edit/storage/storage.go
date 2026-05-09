package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

type Storage struct{ pool *pgxpool.Pool }

func New(pool *pgxpool.Pool) *Storage { return &Storage{pool: pool} }

// Update edits the comment if it belongs to userID, returning the updated row.
// Ownership is enforced via WHERE — non-owners get pgx.ErrNoRows from the
// RETURNING clause, which the usecase maps to errs.ErrNotFound.
func (s *Storage) Update(ctx context.Context, commentID, userID, body string) (*models.Comment, error) {
	var c models.Comment
	err := s.pool.QueryRow(ctx,
		`UPDATE game_comments SET body = $1, updated_at = now()
		 WHERE id = $2 AND user_id = $3
		 RETURNING id, game_id, user_id, body, created_at, COALESCE(updated_at, created_at)`,
		body, commentID, userID,
	).Scan(&c.ID, &c.GameID, &c.UserID, &c.Body, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
