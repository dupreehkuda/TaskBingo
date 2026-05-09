package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

type Storage struct{ pool *pgxpool.Pool }

func New(pool *pgxpool.Pool) *Storage { return &Storage{pool: pool} }

// GetMany returns packs by id, hiding private packs that the requester didn't create.
func (s *Storage) GetMany(ctx context.Context, packIDs []string, requesterID string) (models.Packs, error) {
	rows, err := s.pool.Query(ctx,
		`SELECT id, title, tasks FROM packs
		 WHERE id = ANY($1::uuid[]) AND (is_private = false OR creator = $2)`,
		packIDs, requesterID,
	)
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
