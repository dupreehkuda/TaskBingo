package userRepository

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

func (u userRepository) GetGame(gameID string) (*models.Game, error) {
	game, err := u.conn.GetGame(context.Background(), &api.GetGameRequest{GameID: gameID})
	if err != nil {
		u.logger.Error("Error in call to repository service", zap.Error(err))
		return nil, err
	}

	return &models.Game{
		GameID:       game.GameID,
		User1Id:      game.User1Id,
		User2Id:      game.User2Id,
		PackId:       game.Pack,
		Status:       game.Status,
		User1Bingo:   game.User1Bingo,
		User2Bingo:   game.User2Bingo,
		Winner:       game.Winner,
		Numbers:      game.Numbers,
		User1Numbers: game.User1Numbers,
		User2Numbers: game.User2Numbers,
	}, nil
}
func (u userRepository) AchieveGame(game *models.Game) error {
	_, err := u.conn.AchieveGame(context.Background(), &api.GameRequest{
		GameID:       game.GameID,
		User1Id:      game.User1Id,
		User2Id:      game.User2Id,
		Pack:         game.PackId,
		Status:       game.Status,
		User1Bingo:   game.User1Bingo,
		User2Bingo:   game.User2Bingo,
		Winner:       game.Winner,
		Numbers:      game.Numbers,
		User1Numbers: game.User1Numbers,
		User2Numbers: game.User2Numbers,
	})

	if err != nil {
		u.logger.Error("Error in call to repository service", zap.Error(err))
		return err
	}

	return nil
}
