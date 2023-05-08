package userRepository

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// GetGame calls user service to get current game
func (u userRepository) GetGame(gameID string) (*models.Game, error) {
	game, err := u.conn.GetGame(context.Background(), &api.GetGameRequest{GameID: gameID})
	if err != nil {
		u.logger.Error("Error in call to repository service", zap.Error(err))
		return nil, err
	}

	return mapFromGameRequest(game), nil
}

// AchieveGame calls user service to finish the game
func (u userRepository) AchieveGame(game *models.Game) error {
	_, err := u.conn.AchieveGame(context.Background(), mapToGameRequest(game))

	if err != nil {
		u.logger.Error("Error in call to repository service", zap.Error(err))
		return err
	}

	return nil
}
