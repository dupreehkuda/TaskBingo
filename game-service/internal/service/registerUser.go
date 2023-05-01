package service

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/tokens"
)

// RegisterUser calls user service to register new user and returns JWT-token
func (s service) RegisterUser(credits *models.RegisterCredentials) (string, error) {
	userID, err := s.userRepository.RegisterUser(credits)
	if err != nil {
		return "", err
	}

	token, err := tokens.GenerateJWT(userID, credits.Username)
	if err != nil {
		s.logger.Error("Error while generating jwt", zap.Error(err))
		return "", err
	}

	return token, nil
}
