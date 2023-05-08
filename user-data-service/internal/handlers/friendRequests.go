package handlers

import (
	"context"

	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

// RequestFriend handles the operation of creating a friendship
func (h *Handlers) RequestFriend(ctx context.Context, req *api.FriendRequest) (*api.Empty, error) {
	err := h.service.RequestFriend(ctx, req.UserID, req.FriendID)
	if err != nil {
		h.logger.Error("Error occurred in call to processor", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}

// AcceptFriend handles the operation of accepting a friendship
func (h *Handlers) AcceptFriend(ctx context.Context, req *api.FriendRequest) (*api.Empty, error) {
	err := h.service.AcceptFriend(ctx, req.UserID, req.FriendID)
	if err != nil {
		h.logger.Error("Error occurred in call to processor", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}

// DeleteFriend handles the operation of deleting/canceling friendship
func (h *Handlers) DeleteFriend(ctx context.Context, req *api.FriendRequest) (*api.Empty, error) {
	err := h.service.DeleteFriend(ctx, req.UserID, req.FriendID)
	if err != nil {
		h.logger.Error("Error occurred in call to processor", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}
