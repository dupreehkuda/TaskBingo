package repository

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// CreateGame calls user service to write new game
func (r repository) CreateGame(ctx context.Context, game *models.Game) error {
	_, err := r.conn.CreateGame(ctx, mapToGameRequest(game))

	if err != nil {
		r.logger.Error("Error occurred in connection to user service", zap.Error(err))
		return err
	}

	return nil
}

// AcceptGame calls user service to accept a game
func (r repository) AcceptGame(ctx context.Context, userID, gameID string) error {
	_, err := r.conn.AcceptGame(ctx, mapToStatusGameRequest(userID, gameID))

	if err != nil {
		r.logger.Error("Error occurred in connection to user service", zap.Error(err))
		return err
	}

	return nil
}

// DeleteGame calls user service to delete a game
func (r repository) DeleteGame(ctx context.Context, userID, gameID string) error {
	_, err := r.conn.DeleteGame(ctx, mapToStatusGameRequest(userID, gameID))

	if err != nil {
		r.logger.Error("Error occurred in connection to user service", zap.Error(err))
		return err
	}

	return nil
}
