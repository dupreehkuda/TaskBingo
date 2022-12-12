package handlers

import (
	"context"
	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
	"go.uber.org/zap"
)

func (h *Handlers) GetUserData(ctx context.Context, req *api.GetUserDataRequest) (*api.GetUserDataResponse, error) {
	resp, err := h.processor.GetUserData(req.Login)
	if err != nil {
		h.logger.Error("Unable to call processors", zap.Error(err))
		return nil, err
	}

	return &api.GetUserDataResponse{
		Login:  resp.UserID,
		Points: int32(resp.Points),
		Email:  resp.Email,
	}, nil
}
