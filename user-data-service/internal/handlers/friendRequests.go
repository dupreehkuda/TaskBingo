package handlers

import (
	"context"

	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

func (h *Handlers) RequestFriend(ctx context.Context, req *api.FriendRequest) (*api.Empty, error) {
	err := h.processor.RequestFriend(req.UserID, req.FriendID)
	if err != nil {
		h.logger.Error("Error occurred in call to processor", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}

func (h *Handlers) AcceptFriend(ctx context.Context, req *api.FriendRequest) (*api.Empty, error) {
	err := h.processor.AcceptFriend(req.UserID, req.FriendID)
	if err != nil {
		h.logger.Error("Error occurred in call to processor", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}

func (h *Handlers) DeleteFriend(ctx context.Context, req *api.FriendRequest) (*api.Empty, error) {
	err := h.processor.DeleteFriend(req.UserID, req.FriendID)
	if err != nil {
		h.logger.Error("Error occurred in call to processor", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}
