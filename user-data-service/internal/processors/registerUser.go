package processors

import (
	"go.uber.org/zap"

	errs "github.com/dupreehkuda/TaskBingo/user-data-service/internal/customErrors"
)

// RegisterUser checks user duplicate and registers new user
func (p processor) RegisterUser(login, password, email, city string) error {
	exists, err := p.storage.CheckDuplicateUser(login, email)
	if err != nil {
		p.logger.Error("User check db error", zap.Error(err))
		return err
	}

	if exists {
		return errs.ErrCredentialsInUse
	}

	passwordSalt, err := RandSymbols(10)
	if err != nil {
		p.logger.Error("Generating salt error", zap.Error(err))
		return err
	}

	passwordHash := mdHash(password, passwordSalt)

	err = p.storage.CreateUser(login, email, passwordHash, passwordSalt, city)
	if err != nil {
		p.logger.Error("User creation db error", zap.Error(err))
		return err
	}

	return nil
}
