package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

type Storage struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func New(pool *pgxpool.Pool, logger *zap.Logger) *Storage {
	return &Storage{pool: pool, logger: logger}
}

// Get returns the user account with friends, liked packs, and short games.
// Mirrors legacy user-data-service repository.GetUserData line-for-line in semantics.
func (s *Storage) Get(ctx context.Context, userID string) (*models.UserAccountInfo, error) {
	var (
		out      models.UserAccountInfo
		likedIDs []string
	)

	// Initialise the slice fields so nil slices don't marshal as JSON null.
	// The frontend treats these as arrays unconditionally; getting null back
	// causes runtime crashes for users with no friends/packs/games yet.
	out.Friends = []models.FriendsInfo{}
	out.LikedPacks = []models.TaskPack{}
	out.Games = []models.GameShort{}

	if err := s.pool.QueryRow(ctx,
		`SELECT id, username, city, wins, lose, bingo, solo_bingo, likedpacks, ratedpacks
         FROM users WHERE id = $1`, userID,
	).Scan(&out.UserID, &out.Username, &out.City, &out.Wins, &out.Lose, &out.Bingo, &out.SoloBingo, &likedIDs, &out.RatedPacks); err != nil {
		return nil, err
	}

	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		rows, err := s.pool.Query(egCtx,
			`SELECT friends.friend_id,
                    (SELECT users.username FROM users WHERE users.id = friends.friend_id),
                    friends.status, friends.wins, friends.loses
             FROM friends WHERE id = $1`, userID,
		)
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			var f models.FriendsInfo
			if err := rows.Scan(&f.UserID, &f.Username, &f.Status, &f.Wins, &f.Loses); err != nil {
				return err
			}
			out.Friends = append(out.Friends, f)
		}
		return rows.Err()
	})

	eg.Go(func() error {
		rows, err := s.pool.Query(egCtx,
			`SELECT id, title, tasks FROM packs
			 WHERE id = ANY($1::uuid[]) AND (is_private = false OR creator = $2)`,
			likedIDs, userID,
		)
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			var p models.TaskPack
			if err := rows.Scan(&p.ID, &p.Pack.Title, &p.Pack.Tasks); err != nil {
				return err
			}
			out.LikedPacks = append(out.LikedPacks, p)
		}
		return rows.Err()
	})

	eg.Go(func() error {
		rows, err := s.pool.Query(egCtx,
			`SELECT id, user1_id, COALESCE(user2_id::text, ''), pack_id, status,
			        user1_bingo, user2_bingo, kind
			 FROM games WHERE (user1_id = $1 OR user2_id = $1) AND status != 3 LIMIT 15`, userID,
		)
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			var g models.GameShort
			if err := rows.Scan(&g.GameID, &g.User1Id, &g.User2Id, &g.PackId, &g.Status, &g.User1Bingo, &g.User2Bingo, &g.Kind); err != nil {
				return err
			}
			out.Games = append(out.Games, g)
		}
		return rows.Err()
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return &out, nil
}
