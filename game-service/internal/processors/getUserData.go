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

	resp := &models.UserAccountInfo{
		Login:      userInfo.Login,
		City:       userInfo.City,
		Wins:       userInfo.Wins,
		Lose:       userInfo.Lose,
		Scoreboard: userInfo.Scoreboard,
		Friends:    userInfo.Friends,
		Packs:      []models.TaskPack{},
	}

	if len(userInfo.Packs) != 0 {
		tasks, err := p.taskStorage.GetMultiplePacks(userInfo.Packs)
		if err != nil {
			p.logger.Error("Error occurred in call to task service", zap.Error(err))
			return nil, err
		}

		resp.Packs = *tasks
		return resp, nil
	}

	return resp, nil
}
