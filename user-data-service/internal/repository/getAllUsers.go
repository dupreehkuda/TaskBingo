package repository

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// GetAllUsers retrieves all users from database
func (r repository) GetAllUsers(ctx context.Context) (*[]models.AllUsers, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		r.logger.Error("Error while acquiring connection", zap.Error(err))
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, "SELECT id, username, city, bingo FROM users ORDER BY bingo DESC;")
	if err != nil {
		r.logger.Error("Error while executing query", zap.Error(err))
		return nil, err
	}

	var resp []models.AllUsers

	for rows.Next() {
		var u models.AllUsers
		err = rows.Scan(&u.UserID, &u.Username, &u.City, &u.Bingo)
		if err != nil {
			r.logger.Error("Error while scanning query", zap.Error(err))
			return nil, err
		}

		resp = append(resp, u)
	}

	return &resp, nil
}
