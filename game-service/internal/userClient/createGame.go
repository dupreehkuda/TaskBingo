package userClient

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// CreateGame calls user service to write new game
func (u userClient) CreateGame(game *models.Game) error {
	_, err := u.conn.CreateGame(context.Background(), &api.NewGameRequest{
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
		u.logger.Error("Error occurred in connection to user service", zap.Error(err))
		return err
	}

	return nil
}
