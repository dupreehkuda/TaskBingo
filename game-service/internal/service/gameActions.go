package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// GetGame retrieves game from repository
func (s service) GetGame(ctx context.Context, gameID string) (*models.Game, error) {
	return s.repository.GetGame(ctx, gameID)
}

// GetRoom retrieves game from repository and returns new models.Room
func (s service) GetRoom(ctx context.Context, gameID string) (*models.Room, error) {
	gameInfo, err := s.repository.GetGame(ctx, gameID)
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
func (s service) UpdateGame(ctx context.Context, room *models.Room, action *models.GameAction) (*models.GameUpdate, error) {
	update := &models.GameUpdate{}

	if room.Player1 == nil && room.Player2 == nil {
		return nil, nil
	}

	if room.Status != models.GameInProcess || room.Status != models.GameOneFinished || room.Status != models.GameEnd {
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
		room.Game.User1Numbers, room.Player1.Finished = action.Numbers, action.Finished

		update.UserID = room.Game.User1Id
		update.Numbers = room.Game.User1Numbers

		if newBingo > room.Game.User1Bingo {
			room.Game.User1Bingo, room.Game.User1Numbers = newBingo, action.Numbers
			update.Numbers = room.Game.User1Numbers
		}

	case room.Game.User2Id:
		room.Game.User2Numbers, room.Player2.Finished = action.Numbers, action.Finished

		update.UserID = room.Game.User2Id
		update.Numbers = room.Game.User2Numbers
		update.Bingo = newBingo

		if newBingo > room.Game.User2Bingo {
			room.Game.User2Bingo, room.Game.User2Numbers = newBingo, action.Numbers
			update.Numbers = room.Game.User2Numbers
		}
	}

	update.Status, update.Bingo = formStatus(room), newBingo

	if update.Status == models.GameEnd {
		setWinner(room)
		return s.achieveGame(ctx, room, update)
	}

	return update, nil
}

// achieveGame writes ended game to repository
func (s service) achieveGame(ctx context.Context, room *models.Room, update *models.GameUpdate) (*models.GameUpdate, error) {
	if err := s.repository.AchieveGame(ctx, room.Game); err != nil {
		s.logger.Error("Error in call to user service")
		return nil, err
	}

	return update, nil
}

// formStatus updates room status if anybody finished the game
func formStatus(room *models.Room) int {
	// if both players finished the game
	if room.Player1 != nil && room.Player2 != nil {
		if room.Player1.Finished && room.Player2.Finished {
			room.Status = models.GameEnd
			return room.Status
		}
	}

	// if only one finished the game
	if room.Player1 != nil && room.Player2 != nil {
		if (room.Player1.Finished || room.Player2.Finished) && room.Status == models.GameInProcess {
			room.Status = models.GameOneFinished
			return room.Status
		}
	}

	// if only one returned to the game
	if room.Player1 != nil && room.Player2 != nil {
		if !room.Player1.Finished && !room.Player2.Finished && room.Status == models.GameOneFinished {
			room.Status = models.GameInProcess
			return room.Status
		}
	}

	// if one is not in the game
	if room.Player1 == nil || room.Player2 == nil {
		room.Status = models.GameWaiting
		return room.Status
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
