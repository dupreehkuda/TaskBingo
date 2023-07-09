package service

import (
	"context"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/tokens"
)

// RegisterUser calls user service to register new user and returns JWT-token
func (s service) RegisterUser(ctx context.Context, credits *models.RegisterCredentials) (string, error) {
	userID, err := s.repository.RegisterUser(ctx, credits)
	if err != nil {
		return "", err
	}

	return tokens.GenerateJWT(userID, credits.Username)
}
