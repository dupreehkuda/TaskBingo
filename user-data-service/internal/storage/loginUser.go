package storage

import (
	"context"

	"go.uber.org/zap"

	"github.com/georgysavva/scany/v2/pgxscan"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// LoginUser gets user's data from the database to check for correct credentials
func (s storage) LoginUser(username string) (*models.LoginUserResponse, error) {
	var data models.LoginUserResponse

	const query = "SELECT id, passwordhash, passwordsalt FROM login WHERE id = (SELECT id from users where username = $1);"
	err := pgxscan.Get(context.Background(), s.pool, &data, query, username)
	if err != nil {
		s.logger.Error("Error occurred while getting login data", zap.Error(err))
		return nil, err
	}

	return &data, nil
}
