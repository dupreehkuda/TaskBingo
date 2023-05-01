package handlers

import (
	"context"

	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

// GetUserData handles the operation of getting some important user data for the account
func (h *Handlers) GetUserData(ctx context.Context, req *api.GetUserDataRequest) (*api.GetUserDataResponse, error) {
	resp, err := h.service.GetUserData(req.UserID)
	if err != nil {
		h.logger.Error("Unable to call service", zap.Error(err))
		return nil, err
	}

	ans := &api.GetUserDataResponse{
		UserID:     resp.UserID,
		Username:   resp.Username,
		City:       resp.City,
		Wins:       int32(resp.Wins),
		Loses:      int32(resp.Lose),
		Bingo:      int32(resp.Bingo),
		Friends:    nil,
		LikedPacks: resp.LikedPacks,
		RatedPacks: resp.RatedPacks,
	}

	for _, val := range resp.Friends {
		ans.Friends = append(ans.Friends, &api.FriendInfo{
			UserID:   val.UserID,
			Username: val.Username,
			Status:   int32(val.Status),
			Wins:     int32(val.Wins),
			Loses:    int32(val.Loses),
		})
	}

	return ans, nil
}
