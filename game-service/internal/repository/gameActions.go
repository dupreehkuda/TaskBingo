package repository

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// GetGame calls user service to get current game
func (r repository) GetGame(ctx context.Context, gameID string) (*models.Game, error) {
	game, err := r.conn.GetGame(ctx, &api.GetGameRequest{GameID: gameID})
	if err != nil {
		r.logger.Error("Error in call to repository service", zap.Error(err))
		return nil, err
	}

	return mapFromGameRequest(game), nil
}

// AchieveGame calls user service to finish the game
func (r repository) AchieveGame(ctx context.Context, game *models.Game) error {
	_, err := r.conn.AchieveGame(ctx, mapToGameRequest(game))

	if err != nil {
		r.logger.Error("Error in call to repository service", zap.Error(err))
		return err
	}

	return nil
}
