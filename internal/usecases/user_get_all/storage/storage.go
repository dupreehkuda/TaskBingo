package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

type Storage struct{ pool *pgxpool.Pool }

func New(pool *pgxpool.Pool) *Storage { return &Storage{pool: pool} }

func (s *Storage) ListAll(ctx context.Context) (models.Users, error) {
	rows, err := s.pool.Query(ctx,
		`SELECT id, username, city, bingo FROM users ORDER BY bingo DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out models.Users
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.UserID, &u.Username, &u.City, &u.Bingo); err != nil {
			return nil, err
		}
		out = append(out, u)
	}
	return out, rows.Err()
}
