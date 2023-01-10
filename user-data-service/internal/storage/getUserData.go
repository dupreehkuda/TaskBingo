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
	var friends []string

	row := conn.QueryRow(ctx, "SELECT login, city, wins, lose, scoreboard, friends, packs FROM users WHERE login = $1", login)
	err = row.Scan(&resp.Login, &resp.City, &resp.Wins, &resp.Lose, &resp.Scoreboard, &friends, &resp.Packs)
	if err != nil {
		s.logger.Error("Error when executing statement", zap.Error(err))
		return &resp, err
	}

	rows, err := conn.Query(ctx, "SELECT login, city, scoreboard FROM users WHERE id = ANY($1);", friends)
	if err != nil {
		s.logger.Error("Error when executing statement", zap.Error(err))
		return nil, err
	}

	for rows.Next() {
		var nf models.FriendsInfo
		err = rows.Scan(&nf.Login, &nf.City, &nf.Score)
		if err != nil {
			s.logger.Error("Error when scanning data", zap.Error(err))
			return nil, err
		}

		resp.Friends = append(resp.Friends, nf)
	}

	return &resp, nil
}
