package userRepository

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// CreateGame calls user service to write new game
func (u userRepository) CreateGame(ctx context.Context, game *models.Game) error {
	_, err := u.conn.CreateGame(ctx, mapToGameRequest(game))

	if err != nil {
		u.logger.Error("Error occurred in connection to user service", zap.Error(err))
		return err
	}

	return nil
}

// AcceptGame calls user service to accept a game
func (u userRepository) AcceptGame(ctx context.Context, userID, gameID string) error {
	_, err := u.conn.AcceptGame(ctx, mapToStatusGameRequest(userID, gameID))

	if err != nil {
		u.logger.Error("Error occurred in connection to user service", zap.Error(err))
		return err
	}

	return nil
}

// DeleteGame calls user service to delete a game
func (u userRepository) DeleteGame(ctx context.Context, userID, gameID string) error {
	_, err := u.conn.DeleteGame(ctx, mapToStatusGameRequest(userID, gameID))

	if err != nil {
		u.logger.Error("Error occurred in connection to user service", zap.Error(err))
		return err
	}

	return nil
}
