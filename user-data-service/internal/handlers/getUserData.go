package handlers

import (
	"context"

	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

// GetUserData handles the operation of getting some important user data for the account
func (h *Handlers) GetUserData(ctx context.Context, req *api.GetUserDataRequest) (*api.GetUserDataResponse, error) {
	resp, err := h.service.GetUserData(ctx, req.UserID)
	if err != nil {
		h.logger.Error("Unable to call service", zap.Error(err))
		return nil, err
	}

	return mapToUserDataResponse(resp), nil
}
