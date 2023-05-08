package repository

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// GetUserData retrieves user data from database
func (r repository) GetUserData(ctx context.Context, userID string) (*models.GetUserDataResponse, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		r.logger.Error("Error while acquiring connection", zap.Error(err))
		return &models.GetUserDataResponse{}, err
	}
	defer conn.Release()

	var resp models.GetUserDataResponse

	row := conn.QueryRow(ctx, "SELECT id, username, city, wins, lose, bingo, likedPacks, ratedPacks FROM users WHERE id = $1", userID)
	err = row.Scan(&resp.UserID, &resp.Username, &resp.City, &resp.Wins, &resp.Lose, &resp.Bingo, &resp.LikedPacks, &resp.RatedPacks)
	if err != nil {
		r.logger.Error("Error when executing statement", zap.Error(err))
		return &resp, err
	}

	rows, err := conn.Query(ctx, "SELECT friends.friend_id, (SELECT users.username FROM users WHERE users.id = friends.friend_id) AS username, friends.status, friends.wins, friends.loses FROM friends WHERE id = $1;", userID)
	if err != nil {
		r.logger.Error("Error when executing statement", zap.Error(err))
		return nil, err
	}

	for rows.Next() {
		var nf models.FriendsInfo
		err = rows.Scan(&nf.UserID, &nf.Username, &nf.Status, &nf.Wins, &nf.Loses)
		if err != nil {
			r.logger.Error("Error when scanning data", zap.Error(err))
			return nil, err
		}

		resp.Friends = append(resp.Friends, nf)
	}

	return &resp, nil
}