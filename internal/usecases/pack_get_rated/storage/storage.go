package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

type Storage struct{ pool *pgxpool.Pool }

func New(pool *pgxpool.Pool) *Storage { return &Storage{pool: pool} }

// TopRated returns public packs sorted by rating, plus any private packs the
// requesting user created (so they can show up in their "Mine" tab even
// though they're invisible to everyone else).
func (s *Storage) TopRated(ctx context.Context, userID string) (models.Packs, error) {
	rows, err := s.pool.Query(ctx,
		`SELECT id, title, tasks, is_private, creator::text
		 FROM packs
		 WHERE is_private = false OR creator = $1
		 ORDER BY rating DESC`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := models.Packs{}
	for rows.Next() {
		var p models.TaskPack
		if err := rows.Scan(&p.ID, &p.Pack.Title, &p.Pack.Tasks, &p.IsPrivate, &p.Creator); err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, rows.Err()
}
