package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dupreehkuda/TaskBingo/internal/models"
	"github.com/dupreehkuda/TaskBingo/internal/pkg/errs"
)

type Storage struct{ pool *pgxpool.Pool }

func New(pool *pgxpool.Pool) *Storage { return &Storage{pool: pool} }

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
