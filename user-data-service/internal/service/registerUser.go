package service

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	errs "github.com/dupreehkuda/TaskBingo/user-data-service/internal/customErrors"
)

// RegisterUser checks user duplicate and registers new user
func (s service) RegisterUser(ctx context.Context, username, password, email, city string) (string, error) {
	exists, err := s.repository.CheckDuplicateUser(ctx, username, email)
	if err != nil {
		s.logger.Error("User check db error", zap.Error(err))
		return "", err
	}

	if exists {
		return "", errs.ErrCredentialsInUse
	}

	passwordSalt, err := RandSymbols(10)
	if err != nil {
		s.logger.Error("Generating salt error", zap.Error(err))
		return "", err
	}

	passwordHash := mdHash(password, passwordSalt)
	userID, err := uuid.NewUUID()
	if err != nil {
		s.logger.Error("Generating UUID error", zap.Error(err))
		return "", err
	}

	err = s.repository.CreateUser(ctx, userID.String(), username, email, passwordHash, passwordSalt, cases.Title(language.English).String(city))
	if err != nil {
		s.logger.Error("User creation db error", zap.Error(err))
		return "", err
	}

	return userID.String(), nil
}
