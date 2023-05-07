package service

import (
	"math/rand"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// CreateGame creates new game instance
func (s service) CreateGame(userID, opponentID, packID string) error {
	var newGame = models.Game{
		GameID:       uuid.New().String(),
		User1Id:      userID,
		User2Id:      opponentID,
		PackId:       packID,
		Status:       0,
		User1Bingo:   0,
		User2Bingo:   0,
		Winner:       "",
		Numbers:      newShuffledNumberSet(),
		User1Numbers: newDefaultNumberSet(),
		User2Numbers: newDefaultNumberSet(),
	}

	if err := s.userRepository.CreateGame(&newGame); err != nil {
		s.logger.Error("Error occurred in call to user service", zap.Error(err))
		return err
	}

	return nil
}

// AcceptGame changes status when user accepts the game
func (s service) AcceptGame(userID, gameID string) error {
	if err := s.userRepository.AcceptGame(userID, gameID); err != nil {
		s.logger.Error("Error occurred in call to user service", zap.Error(err))
		return err
	}

	return nil
}

// DeleteGame deletes the game if created incorrectly or declined by user
func (s service) DeleteGame(userID, gameID string) error {
	if err := s.userRepository.DeleteGame(userID, gameID); err != nil {
		s.logger.Error("Error occurred in call to user service", zap.Error(err))
		return err
	}

	return nil
}

// newShuffledNumberSet returns a slice of shuffled numbers 1-16
func newShuffledNumberSet() []int32 {
	numbers := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	rand.Shuffle(len(numbers), func(i, j int) { numbers[i], numbers[j] = numbers[j], numbers[i] })
	return numbers
}

// newShuffledNumberSet returns a slice of 16 zeroes (default user state)
func newDefaultNumberSet() []int32 {
	return make([]int32, 16, 16)
}
