package processors

import (
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// Enumeration for game status
const (
	_ = iota
	Requested
	Started
	Ended
)

// CreateGame creates new game instance
func (p processor) CreateGame(user1, user2, packID string) error {
	var newGame = models.Game{
		GameID:       uuid.New(),
		User1Id:      user1,
		User2Id:      user2,
		PackId:       packID,
		Status:       Requested,
		User1Bingo:   0,
		User2Bingo:   0,
		Winner:       "",
		Numbers:      nil,
		User1Numbers: nil,
		User2Numbers: nil,
	}

	if err := p.userStorage.CreateGame(&newGame); err != nil {
		p.logger.Error("Error occurred in call to user service", zap.Error(err))
		return err
	}

	return nil
}
