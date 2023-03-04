package processors

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// GetUserData gets user's most important info
func (p processor) GetUserData(userID string) (*models.UserAccountInfo, error) {
	userInfo, err := p.userStorage.GetUserData(userID)
	if err != nil {
		p.logger.Error("Error occurred in call to user service", zap.Error(err))
		return nil, err
	}

	resp := &models.UserAccountInfo{
		UserID:     userInfo.UserID,
		Username:   userInfo.Username,
		City:       userInfo.City,
		Wins:       userInfo.Wins,
		Lose:       userInfo.Lose,
		Bingo:      userInfo.Bingo,
		Friends:    userInfo.Friends,
		LikedPacks: []models.TaskPack{},
		RatedPacks: userInfo.RatedPacks,
	}

	if len(userInfo.LikedPacks) != 0 {
		tasks, err := p.taskStorage.GetMultiplePacks(userInfo.LikedPacks)
		if err != nil {
			p.logger.Error("Error occurred in call to task service", zap.Error(err))
			return nil, err
		}

		resp.LikedPacks = *tasks
		return resp, nil
	}

	if len(userInfo.RatedPacks) == 0 {
		resp.RatedPacks = []string{}
	}

	return resp, nil
}
