package userRepository

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// GetGame calls user service to get current game
func (u userRepository) GetGame(ctx context.Context, gameID string) (*models.Game, error) {
	game, err := u.conn.GetGame(ctx, &api.GetGameRequest{GameID: gameID})
	if err != nil {
		u.logger.Error("Error in call to repository service", zap.Error(err))
		return nil, err
	}

	return mapFromGameRequest(game), nil
}

// AchieveGame calls user service to finish the game
func (u userRepository) AchieveGame(ctx context.Context, game *models.Game) error {
	_, err := u.conn.AchieveGame(ctx, mapToGameRequest(game))

	if err != nil {
		u.logger.Error("Error in call to repository service", zap.Error(err))
		return err
	}

	return nil
}
