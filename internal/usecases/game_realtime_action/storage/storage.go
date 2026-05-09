package storage

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/errs"
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

// Get is duplicated with game_get/storage by design (per-usecase storage rule).
func (s *Storage) Get(ctx context.Context, gameID string) (*models.Game, error) {
	var g models.Game
	err := s.pool.QueryRow(ctx,
		`SELECT id, user1_id, COALESCE(user2_id::text, ''), pack_id, status,
                user1_bingo, user2_bingo, COALESCE(winner::text, ''),
                numbers, user1_numbers, user2_numbers, kind
         FROM games WHERE id = $1`, gameID,
	).Scan(&g.GameID, &g.User1Id, &g.User2Id, &g.PackId, &g.Status,
		&g.User1Bingo, &g.User2Bingo, &g.Winner,
		&g.Numbers, &g.User1Numbers, &g.User2Numbers, &g.Kind)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.ErrNotFound
		}
		return nil, err
	}
	return &g, nil
}

// Achieve persists the final game state and bumps both users' bingo, plus
// wins/lose for the winner/loser. Caller must wrap in pgxtx.RunInTx.
func (s *Storage) Achieve(ctx context.Context, g *models.Game) error {
	if _, err := s.q(ctx).Exec(ctx,
		`UPDATE games
         SET status = $1, winner = NULLIF($2, '')::uuid,
             user1_bingo = $3, user2_bingo = $4,
             user1_numbers = $5, user2_numbers = $6,
             finished = $7
         WHERE id = $8`,
		models.GameDBEnded, g.Winner, g.User1Bingo, g.User2Bingo,
		g.User1Numbers, g.User2Numbers, time.Now(), g.GameID,
	); err != nil {
		return err
	}

	if _, err := s.q(ctx).Exec(ctx,
		`UPDATE users SET bingo = bingo + $1 WHERE id = $2`,
		g.User1Bingo, g.User1Id,
	); err != nil {
		return err
	}
	if _, err := s.q(ctx).Exec(ctx,
		`UPDATE users SET bingo = bingo + $1 WHERE id = $2`,
		g.User2Bingo, g.User2Id,
	); err != nil {
		return err
	}

	if g.Winner == g.User1Id {
		if _, err := s.q(ctx).Exec(ctx,
			`UPDATE users SET wins = wins + 1 WHERE id = $1`, g.User1Id,
		); err != nil {
			return err
		}
		_, err := s.q(ctx).Exec(ctx,
			`UPDATE users SET lose = lose + 1 WHERE id = $1`, g.User2Id,
		)
		return err
	}
	if g.Winner == g.User2Id {
		if _, err := s.q(ctx).Exec(ctx,
			`UPDATE users SET wins = wins + 1 WHERE id = $1`, g.User2Id,
		); err != nil {
			return err
		}
		_, err := s.q(ctx).Exec(ctx,
			`UPDATE users SET lose = lose + 1 WHERE id = $1`, g.User1Id,
		)
		return err
	}
	// Tie: bingos already credited, no win/lose increment.
	return nil
}
