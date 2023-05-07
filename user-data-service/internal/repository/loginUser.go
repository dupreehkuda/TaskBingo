package repository

import (
	"context"

	"go.uber.org/zap"

	"github.com/georgysavva/scany/v2/pgxscan"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// LoginUser gets user's data from the database to check for correct credentials
func (r repository) LoginUser(ctx context.Context, username string) (*models.LoginUserResponse, error) {
	var data models.LoginUserResponse

	const query = "SELECT id, passwordhash, passwordsalt FROM login WHERE id = (SELECT id from users where username = $1);"
	err := pgxscan.Get(ctx, r.pool, &data, query, username)
	if err != nil {
		r.logger.Error("Error occurred while getting login data", zap.Error(err))
		return nil, err
	}

	return &data, nil
}
