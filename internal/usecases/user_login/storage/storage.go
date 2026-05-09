package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dupreehkuda/TaskBingo/internal/pkg/errs"
)

type Storage struct{ pool *pgxpool.Pool }

func New(pool *pgxpool.Pool) *Storage { return &Storage{pool: pool} }

type Credentials struct {
	UserID       string
	PasswordHash string
	PasswordSalt string
}

// FetchByUsername returns the login row joined to the user. Returns errs.ErrNotFound if username is unknown.
func (s *Storage) FetchByUsername(ctx context.Context, username string) (Credentials, error) {
	var c Credentials
	err := s.pool.QueryRow(ctx,
		`SELECT id, passwordhash, passwordsalt
         FROM login
         WHERE id = (SELECT id FROM users WHERE username = $1)`,
		username,
	).Scan(&c.UserID, &c.PasswordHash, &c.PasswordSalt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Credentials{}, errs.ErrNotFound
		}
		return Credentials{}, err
	}
	return c, nil
}
