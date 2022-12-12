package processors

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/tokens"
)

// LoginUser calls user service to login user and returns JWT-token
func (p processor) LoginUser(login, password string) (string, error) {
	err := p.userStorage.LoginUser(login, password)
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
