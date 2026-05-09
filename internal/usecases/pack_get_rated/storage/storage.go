package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

type Storage struct{ pool *pgxpool.Pool }

func New(pool *pgxpool.Pool) *Storage { return &Storage{pool: pool} }

func (s *Storage) TopRated(ctx context.Context) (models.Packs, error) {
	rows, err := s.pool.Query(ctx,
		`SELECT id, title, tasks FROM packs WHERE is_private = false ORDER BY rating DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out models.Packs
	for rows.Next() {
		var p models.TaskPack
		if err := rows.Scan(&p.ID, &p.Pack.Title, &p.Pack.Tasks); err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, rows.Err()
}
