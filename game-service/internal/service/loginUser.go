package service

import (
	"context"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/tokens"
)

// LoginUser calls user service to log in user and returns JWT-token
func (s service) LoginUser(ctx context.Context, username, password string) (string, error) {
	userID, err := s.repository.LoginUser(ctx, username, password)
	if err != nil {
		return "", err
	}

	return tokens.GenerateJWT(userID, username)
}
