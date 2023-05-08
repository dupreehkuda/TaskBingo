package service

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// GetRoom retrieves game from repository and returns new models.Room
func (s service) GetRoom(gameID string) (*models.Room, error) {
	gameInfo, err := s.userRepository.GetGame(gameID)
	if err != nil {
		s.logger.Error("Error getting game", zap.Error(err))
		return nil, err
	}

	game := &models.Room{
		Id:     gameID,
		Game:   gameInfo,
		Status: models.GameCreated,
	}

	return game, nil
}

// UpdateGame forms an update on every game event
func (s service) UpdateGame(room *models.Room, action *models.GameAction) (*models.GameUpdate, error) {
	update := &models.GameUpdate{}

	if room.Player1 == nil && room.Player2 == nil {
		return nil, nil
	}

	if room.Status != models.GameInProcess {
		if room.Status == models.GameCreated && (room.Player1 != nil || room.Player2 != nil) {
			room.Status = models.GameWaiting
			return nil, nil
		}

		if room.Status == models.GameWaiting && room.Player1 != nil && room.Player2 != nil {
			room.Status = models.GameStart
			update.Status = room.Status
			return update, nil
		}

		if room.Status == models.GameStart {
			room.Status = models.GameInProcess
		}
	}

	newBingo := countBingo(action.Numbers)

	switch action.UserID {
	case room.Game.User1Id:
		room.Game.User1Numbers = action.Numbers

		update.UserID = room.Game.User1Id
		update.Numbers = room.Game.User1Numbers

		if newBingo > room.Game.User1Bingo {
			room.Game.User1Bingo, room.Game.User1Numbers = newBingo, action.Numbers
			update.Numbers = room.Game.User1Numbers
		}

	case room.Game.User2Id:
		room.Game.User2Numbers = action.Numbers

		update.UserID = room.Game.User2Id
		update.Numbers = room.Game.User2Numbers

		if newBingo > room.Game.User2Bingo {
			room.Game.User2Bingo, room.Game.User2Numbers = newBingo, action.Numbers
			update.Numbers = room.Game.User2Numbers
		}
	}

	update.Status = formStatus(room, action.Finished)

	if update.Status == models.GameEnd {
		setWinner(room)
		return s.achieveGame(room, update)
	}

	return update, nil
}

// achieveGame writes ended game to repository
func (s service) achieveGame(room *models.Room, update *models.GameUpdate) (*models.GameUpdate, error) {
	if err := s.userRepository.AchieveGame(room.Game); err != nil {
		s.logger.Error("Error in call to user service")
		return nil, err
	}

	return update, nil
}

// formStatus updates room status if anybody finished the game
func formStatus(room *models.Room, finished bool) int {
	if finished && room.Status == models.GameInProcess {
		room.Status = models.GameOneFinished
	}

	if finished && room.Status == models.GameOneFinished {
		room.Status = models.GameEnd
	}

	return room.Status
}

// setWinner picks the winner
func setWinner(room *models.Room) {
	if room.Game.User1Bingo > room.Game.User2Bingo {
		room.Game.Winner = room.Game.User1Id
	}

	if room.Game.User2Bingo > room.Game.User1Bingo {
		room.Game.Winner = room.Game.User2Id
	}
}

// countBingo counts bingo amount in provided number set
func countBingo(numbers []int32) int32 {
	var bingos int32

	for i := 0; i < 4; i++ {
		if numbers[i] != 0 && numbers[i+4] != 0 && numbers[i+8] != 0 && numbers[i+12] != 0 {
			bingos += 1
		}
	}

	for i := 0; i < len(numbers); i += 4 {
		if numbers[i] != 0 && numbers[i+1] != 0 && numbers[i+2] != 0 && numbers[i+3] != 0 {
			bingos += 1
		}
	}

	if numbers[0] != 0 && numbers[5] != 0 && numbers[10] != 0 && numbers[15] != 0 {
		bingos += 1
	}

	if numbers[3] != 0 && numbers[6] != 0 && numbers[9] != 0 && numbers[12] != 0 {
		bingos += 1
	}

	return bingos
}
