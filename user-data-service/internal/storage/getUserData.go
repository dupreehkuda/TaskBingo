package storage

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// GetUserData retrieves user data from database
func (s storage) GetUserData(login string) (*models.GetUserDataResponse, error) {
	ctx := context.Background()
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return &models.GetUserDataResponse{}, err
	}
	defer conn.Release()

	var resp models.GetUserDataResponse

	row := conn.QueryRow(ctx, "SELECT login, city, wins, lose, bingo, likedPacks, ratedPacks FROM users WHERE login = $1", login)
	err = row.Scan(&resp.Login, &resp.City, &resp.Wins, &resp.Lose, &resp.Bingo, &resp.LikedPacks, &resp.RatedPacks)
	if err != nil {
		s.logger.Error("Error when executing statement", zap.Error(err))
		return &resp, err
	}

	rows, err := conn.Query(ctx, "SELECT friend, status, wins, loses FROM friends where id = $1;", login)
	if err != nil {
		s.logger.Error("Error when executing statement", zap.Error(err))
		return nil, err
	}

	for rows.Next() {
		var nf models.FriendsInfo
		err = rows.Scan(&nf.Login, &nf.Status, &nf.Wins, &nf.Loses)
		if err != nil {
			s.logger.Error("Error when scanning data", zap.Error(err))
			return nil, err
		}

		resp.Friends = append(resp.Friends, nf)
	}

	return &resp, nil
}
