package service

import (
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// GetUserData gets user's most important info
func (s service) GetUserData(userID string) (*models.UserAccountInfo, error) {
	userInfo, err := s.userRepository.GetUserData(userID)
	if err != nil {
		s.logger.Error("Error occurred in call to user service", zap.Error(err))
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

	if len(userInfo.RatedPacks) == 0 {
		resp.RatedPacks = []string{}
	}

	if len(userInfo.LikedPacks) != 0 {
		tasks, err := s.taskRepository.GetMultiplePacks(userInfo.LikedPacks)
		if err != nil {
			s.logger.Error("Error occurred in call to task service", zap.Error(err))
			return nil, err
		}

		resp.LikedPacks = *tasks
		return resp, nil
	}

	return resp, nil
}
