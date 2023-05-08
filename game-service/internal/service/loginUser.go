package service

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/tokens"
)

// LoginUser calls user service to log in user and returns JWT-token
func (s service) LoginUser(username, password string) (string, error) {
	userID, err := s.userRepository.LoginUser(username, password)
	if err != nil {
		return "", err
	}

	token, err := tokens.GenerateJWT(userID, username)
	if err != nil {
		s.logger.Error("Error while generating jwt", zap.Error(err))
		return "", err
	}

	return token, nil
}