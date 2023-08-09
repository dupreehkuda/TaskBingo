package repository

import (
	"context"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// GetUserData retrieves user data from database
func (r repository) GetUserData(ctx context.Context, userID string) (*models.GetUserDataResponse, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		r.logger.Error("Error while acquiring connection", zap.Error(err))
		return &models.GetUserDataResponse{}, err
	}

	var (
		resp       models.GetUserDataResponse
		likedPacks []string
	)

	row := conn.QueryRow(ctx, "SELECT id, username, city, wins, lose, bingo, likedPacks, ratedPacks FROM users WHERE id = $1", userID)
	err = row.Scan(&resp.UserID, &resp.Username, &resp.City, &resp.Wins, &resp.Lose, &resp.Bingo, &likedPacks, &resp.RatedPacks)
	if err != nil {
		r.logger.Error("Error when executing statement", zap.Error(err))
		return nil, err
	}

	conn.Release()

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		newConn, err := r.pool.Acquire(ctx)
		if err != nil {
			r.logger.Error("Error while acquiring connection", zap.Error(err))
			return err
		}
		defer newConn.Release()

		rows, err := newConn.Query(ctx, "SELECT friends.friend_id, (SELECT users.username FROM users WHERE users.id = friends.friend_id) AS username, friends.status, friends.wins, friends.loses FROM friends WHERE id = $1;", userID)
		if err != nil {
			r.logger.Error("Error when executing statement", zap.Error(err))
			return err
		}

		for rows.Next() {
			var nf models.FriendsInfo

			err = rows.Scan(&nf.UserID, &nf.Username, &nf.Status, &nf.Wins, &nf.Loses)
			if err != nil {
				r.logger.Error("Error when scanning data", zap.Error(err))
				return err
			}

			resp.Friends = append(resp.Friends, nf)
		}

		return nil
	})

	eg.Go(func() error {
		newConn, err := r.pool.Acquire(ctx)
		if err != nil {
			r.logger.Error("Error while acquiring connection", zap.Error(err))
			return err
		}
		defer newConn.Release()

		rows, err := newConn.Query(ctx, "SELECT id, title, tasks FROM packs WHERE id = ANY($1::uuid[])", likedPacks)
		if err != nil {
			r.logger.Error("Error when executing statement", zap.Error(err))
			return err
		}

		for rows.Next() {
			var np models.TaskPack

			err = rows.Scan(&np.ID, &np.Pack.Title, &np.Pack.Tasks)
			if err != nil {
				r.logger.Error("Error when scanning data", zap.Error(err))
				return err
			}

			resp.LikedPacks = append(resp.LikedPacks, np)
		}

		return nil
	})

	eg.Go(func() error {
		newConn, err := r.pool.Acquire(ctx)
		if err != nil {
			r.logger.Error("Error while acquiring connection", zap.Error(err))
			return err
		}
		defer newConn.Release()

		rows, err := newConn.Query(ctx, "SELECT id, user1_id, user2_id, pack_id, status, user1_bingo, user2_bingo FROM games WHERE (user1_id = $1 OR user2_id = $1) AND status != 3 LIMIT 15;", userID)
		if err != nil {
			r.logger.Error("Error when executing statement", zap.Error(err))
			return err
		}

		for rows.Next() {
			var ng models.GameShort

			err = rows.Scan(&ng.GameID, &ng.User1Id, &ng.User2Id, &ng.PackId, &ng.Status, &ng.User1Bingo, &ng.User2Bingo)
			if err != nil {
				r.logger.Error("Error when scanning data", zap.Error(err))
				return err
			}

			resp.GamesShort = append(resp.GamesShort, ng)
		}

		return nil
	})

	if err = eg.Wait(); err != nil {
		r.logger.Error("Error when executing statement in eg", zap.Error(err))
		return nil, err
	}

	return &resp, nil
}
