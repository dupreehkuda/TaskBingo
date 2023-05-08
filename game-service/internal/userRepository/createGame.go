package userRepository

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// CreateGame calls user service to write new game
func (u userRepository) CreateGame(game *models.Game) error {
	_, err := u.conn.CreateGame(context.Background(), &api.GameRequest{
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

// AcceptGame calls user service to accept a game
func (u userRepository) AcceptGame(userID, gameID string) error {
	_, err := u.conn.AcceptGame(context.Background(), &api.StatusGameRequest{
		UserID: userID,
		GameID: gameID,
	})

	if err != nil {
		u.logger.Error("Error occurred in connection to user service", zap.Error(err))
		return err
	}

	return nil
}

// DeleteGame calls user service to delete a game
func (u userRepository) DeleteGame(userID, gameID string) error {
	_, err := u.conn.DeleteGame(context.Background(), &api.StatusGameRequest{
		UserID: userID,
		GameID: gameID,
	})

	if err != nil {
		u.logger.Error("Error occurred in connection to user service", zap.Error(err))
		return err
	}

	return nil
}
