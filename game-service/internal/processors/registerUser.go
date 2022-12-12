package processors

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/tokens"
)

// RegisterUser calls user service to register new user and returns JWT-token
func (p processor) RegisterUser(login, email, password string) (string, error) {
	err := p.userStorage.RegisterUser(login, email, password)
	if err != nil {
		return "", err
	}

	token, err := tokens.GenerateJWT(login)
	if err != nil {
		p.logger.Error("Error while generating jwt", zap.Error(err))
		return "", err
	}

	return token, nil
}
