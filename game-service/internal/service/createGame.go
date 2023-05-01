package service

import (
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// Enum for game status
const (
	_ = iota
	Requested
	Started
	Ended
)

// CreateGame creates new game instance
func (s service) CreateGame(userID, opponentID, packID string) error {
	var newGame = models.Game{
		GameID:       uuid.New().String(),
		User1Id:      userID,
		User2Id:      opponentID,
		PackId:       packID,
		Status:       Requested,
		User1Bingo:   0,
		User2Bingo:   0,
		Winner:       "",
		Numbers:      nil,
		User1Numbers: nil,
		User2Numbers: nil,
	}

	if err := s.userRepository.CreateGame(&newGame); err != nil {
		s.logger.Error("Error occurred in call to user service", zap.Error(err))
		return err
	}

	return nil
}
