package processors

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/tokens"
)

// LoginUser calls user service to log in user and returns JWT-token
func (p processor) LoginUser(username, password string) (string, error) {
	userID, err := p.userStorage.LoginUser(username, password)
	if err != nil {
		return "", err
	}

	token, err := tokens.GenerateJWT(userID, username)
	if err != nil {
		p.logger.Error("Error while generating jwt", zap.Error(err))
		return "", err
	}

	return token, nil
}
