package storage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/pkg/pgxtx"
)

type Querier interface {
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

type Storage struct{ pool *pgxpool.Pool }

func New(pool *pgxpool.Pool) *Storage { return &Storage{pool: pool} }

func (s *Storage) q(ctx context.Context) Querier {
	if tx, ok := pgxtx.From(ctx); ok {
		return tx
	}
	return s.pool
}

// Finalize persists the final solo state and bumps users.solo_bingo. Caller
// must wrap in pgxtx.RunInTx. Returns the games-update affected-rows count so
// the caller can detect a wrong owner / non-solo / unknown game with one read.
func (s *Storage) Finalize(ctx context.Context, userID, gameID string, userNumbers []int32, bingo int32) (int64, error) {
	tag, err := s.q(ctx).Exec(ctx,
		`UPDATE games
		 SET status = $1, user1_numbers = $2, user1_bingo = $3, finished = $4
		 WHERE id = $5 AND user1_id = $6 AND kind = 'solo'`,
		models.GameDBEnded, userNumbers, bingo, time.Now(), gameID, userID,
	)
	if err != nil {
		return 0, err
	}
	if tag.RowsAffected() == 0 {
		return 0, nil
	}

	if _, err := s.q(ctx).Exec(ctx,
		`UPDATE users SET solo_bingo = solo_bingo + $1 WHERE id = $2`,
		bingo, userID,
	); err != nil {
		return tag.RowsAffected(), err
	}
	return tag.RowsAffected(), nil
}
