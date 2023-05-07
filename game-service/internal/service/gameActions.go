package service

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

const (
	_ = iota
	Created
	Waiting
	Start
	InProcess
	OneFinished
	End
)

func (s service) GetRoom(gameID string) (*models.Room, error) {
	gameInfo, err := s.userRepository.GetGame(gameID)
	if err != nil {
		s.logger.Error("Error getting game", zap.Error(err))
		return nil, err
	}

	game := &models.Room{
		Id:     gameID,
		Game:   gameInfo,
		Status: Created,
	}

	return game, nil
}

func (s service) UpdateGame(room *models.Room, action *models.GameAction) (*models.GameUpdate, error) {
	fmt.Println("before changing inside func", &room)
	update := &models.GameUpdate{}

	if room.Player1 == nil && room.Player2 == nil {
		s.logger.Debug("hren kakayato")
		return nil, nil
	}

	if room.Status != InProcess {
		if room.Status == Created && room.Player1 != nil || room.Player2 != nil {
			s.logger.Debug("status changed to waiting")
			room.Status = Waiting
			fmt.Println("after changing inside func", &room)
			return nil, nil
		}

		if room.Status == Waiting && room.Player1 != nil && room.Player2 != nil {
			s.logger.Debug("status changed to start")
			room.Status = Start
			update.Status = room.Status
			return update, nil
		}

		if room.Status == Start {
			s.logger.Debug("status changed to inprocess")
			room.Status = InProcess
		}
	}

	newBingo := countBingo(action.Numbers)

	switch action.UserID {
	case room.Game.User1Id:
		update.UserID = room.Game.User2Id

		if newBingo > room.Game.User1Bingo {
			room.Game.User1Bingo, room.Game.User1Numbers = newBingo, action.Numbers
			update.Numbers = room.Game.User1Numbers
		}

	case room.Game.User2Id:
		update.UserID = room.Game.User1Id

		if newBingo > room.Game.User2Bingo {
			room.Game.User2Bingo, room.Game.User2Numbers = newBingo, action.Numbers
			update.Numbers = room.Game.User1Numbers
		}
	}

	room.Status = formStatus(action.Finished, room.Status)
	update.Status = room.Status

	if update.Status == End {
		return s.achieveGame(room, update)
	}

	return update, nil
}

func (s service) achieveGame(room *models.Room, update *models.GameUpdate) (*models.GameUpdate, error) {
	if err := s.userRepository.AchieveGame(room.Game); err != nil {
		s.logger.Error("Error in call to user service")
		return nil, err
	}

	return update, nil
}

func formStatus(finished bool, roomStatus int) int {
	if finished && roomStatus == InProcess {
		return OneFinished
	}

	if finished && roomStatus == OneFinished {
		return End
	}

	return roomStatus
}

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
