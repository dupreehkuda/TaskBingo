package processors

import (
	"go.uber.org/zap"

	errs "github.com/dupreehkuda/TaskBingo/user-data-service/internal/customErrors"
)

func (p processor) LoginUser(login, password string) error {
	resp, err := p.storage.LoginUser(login)
	if err != nil {
		p.logger.Error("Error occurred in call to storage", zap.Error(err))
		return errs.ErrWrongCredentials
	}

	if resp == nil {
		p.logger.Error("Something went wrong. resp == nil")
		return errs.ErrWrongCredentials
	}

	checkHash := mdHash(password, resp.PasswordSalt)
	if checkHash != resp.PasswordHash {
		return errs.ErrWrongCredentials
	}

	return nil
}
