package storage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/usecases/user_stats/entity"
)

type Storage struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func New(pool *pgxpool.Pool, logger *zap.Logger) *Storage {
	return &Storage{pool: pool, logger: logger}
}

// FetchFinishedGames returns finished games for the user from `from` (inclusive)
// onwards, joined with the pack title. Caller is responsible for time bucketing
// and aggregation. Both user1 and user2 sides are returned in a single sweep so
// the indexed lookup can run via UNION ALL.
func (s *Storage) FetchFinishedGames(ctx context.Context, userID string, from time.Time) ([]entity.GameRow, error) {
	const q = `
		SELECT g.id, g.kind, g.finished, g.pack_id,
		       COALESCE(g.winner::text, ''),
		       g.user1_id, COALESCE(g.user2_id::text, ''),
		       g.user1_bingo, g.user2_bingo,
		       g.user1_numbers, g.user2_numbers,
		       p.title
		  FROM games g
		  JOIN packs p ON p.id = g.pack_id
		 WHERE (g.user1_id = $1 OR g.user2_id = $1)
		   AND g.status = 3
		   AND g.finished >= $2
		 ORDER BY g.finished ASC
	`
	rows, err := s.pool.Query(ctx, q, userID, from)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]entity.GameRow, 0, 32)
	for rows.Next() {
		var g entity.GameRow
		if err := rows.Scan(
			&g.GameID, &g.Kind, &g.Finished, &g.PackID,
			&g.Winner,
			&g.User1ID, &g.User2ID,
			&g.User1Bingo, &g.User2Bingo,
			&g.User1Numbers, &g.User2Numbers,
			&g.PackTitle,
		); err != nil {
			return nil, err
		}
		out = append(out, g)
	}
	return out, rows.Err()
}
