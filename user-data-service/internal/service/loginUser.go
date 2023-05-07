package service

import (
	"context"

	"go.uber.org/zap"

	errs "github.com/dupreehkuda/TaskBingo/user-data-service/internal/customErrors"
)

// LoginUser checks provided credentials and logs user
func (s service) LoginUser(ctx context.Context, username, password string) (string, error) {
	resp, err := s.repository.LoginUser(ctx, username)
	if err != nil {
		s.logger.Error("Error occurred in call to repository", zap.Error(err))
		return "", errs.ErrWrongCredentials
	}

	if resp == nil {
		s.logger.Error("Something went wrong. resp == nil")
		return "", errs.ErrWrongCredentials
	}

	checkHash := mdHash(password, resp.PasswordSalt)
	if checkHash != resp.PasswordHash {
		return "", errs.ErrWrongCredentials
	}

	return resp.UserID, nil
}
