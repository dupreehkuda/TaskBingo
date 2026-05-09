package storage

import (
	"context"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dupreehkuda/TaskBingo/pkg/pgxtx"
)

type Querier interface {
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type Storage struct{ pool *pgxpool.Pool }

func New(pool *pgxpool.Pool) *Storage { return &Storage{pool: pool} }

func (s *Storage) q(ctx context.Context) Querier {
	if tx, ok := pgxtx.From(ctx); ok {
		return tx
	}
	return s.pool
}

// CheckDuplicate returns true if username or email already exists.
func (s *Storage) CheckDuplicate(ctx context.Context, username, email string) (bool, error) {
	var exists bool
	err := s.q(ctx).QueryRow(ctx,
		`SELECT EXISTS (SELECT 1 FROM users WHERE username = $1 OR email = $2)`,
		username, email,
	).Scan(&exists)
	return exists, err
}

// Insert writes the user row plus the login row in the active tx (caller must wrap in pgxtx.RunInTx).
func (s *Storage) Insert(ctx context.Context, userID, username, email, city, passwordHash, passwordSalt string) error {
	if _, err := s.q(ctx).Exec(ctx,
		`INSERT INTO users (id, username, email, registered, city) VALUES ($1, $2, $3, $4, $5)`,
		userID, strings.TrimSpace(username), strings.TrimSpace(email), time.Now(), strings.TrimSpace(city),
	); err != nil {
		return err
	}
	_, err := s.q(ctx).Exec(ctx,
		`INSERT INTO login (id, passwordhash, passwordsalt) VALUES ($1, $2, $3)`,
		userID, passwordHash, passwordSalt,
	)
	return err
}
