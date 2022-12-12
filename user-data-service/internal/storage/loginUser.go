package storage

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
	"github.com/georgysavva/scany/v2/pgxscan"
)

// LoginUser gets user's data from the database to check for correct credentials
func (s storage) LoginUser(login string) (*models.LoginUserResponse, error) {
	var data models.LoginUserResponse

	const query = "SELECT passwordhash, passwordsalt FROM login WHERE id = $1;"

	err := pgxscan.Get(context.Background(), s.pool, &data, query, login)
	if err != nil {
		s.logger.Error("Error occurred while getting login data", zap.Error(err))
		return nil, err
	}

	return &data, nil
}
