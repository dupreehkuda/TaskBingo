package processors

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// GetUserData gets user's most important info
func (p processor) GetUserData(login string) (*models.UserAccountInfo, error) {
	userInfo, err := p.userStorage.GetUserData(login)
	if err != nil {
		p.logger.Error("Error occurred in call to user service", zap.Error(err))
		return nil, err
	}

	tasks, err := p.taskStorage.GetFavouritePacks(userInfo.Packs)
	if err != nil {
		p.logger.Error("Error occurred in call to task service", zap.Error(err))
		return nil, err
	}

	return &models.UserAccountInfo{
		Login:      userInfo.Login,
		City:       userInfo.City,
		Wins:       userInfo.Wins,
		Lose:       userInfo.Lose,
		Scoreboard: userInfo.Scoreboard,
		Friends:    userInfo.Friends,
		Packs:      *tasks,
	}, nil
}
