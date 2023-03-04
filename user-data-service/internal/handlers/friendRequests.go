package handlers

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

func (h *Handlers) RequestFriend(ctx context.Context, req *api.FriendRequest) (*api.Empty, error) {
	userID, err := uuid.Parse(req.UserID.Id)
	if err != nil {
		h.logger.Error("Unable to parse uuid", zap.Error(err))
		return nil, err
	}

	friendID, err := uuid.Parse(req.FriendID.Id)
	if err != nil {
		h.logger.Error("Unable to parse uuid", zap.Error(err))
		return nil, err
	}

	err = h.processor.RequestFriend(userID, friendID)
	if err != nil {
		h.logger.Error("Error occurred in call to processor", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}

func (h *Handlers) AcceptFriend(ctx context.Context, req *api.FriendRequest) (*api.Empty, error) {
	userID, err := uuid.Parse(req.UserID.Id)
	if err != nil {
		h.logger.Error("Unable to parse uuid", zap.Error(err))
		return nil, err
	}

	friendID, err := uuid.Parse(req.FriendID.Id)
	if err != nil {
		h.logger.Error("Unable to parse uuid", zap.Error(err))
		return nil, err
	}

	err = h.processor.AcceptFriend(userID, friendID)
	if err != nil {
		h.logger.Error("Error occurred in call to processor", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}

func (h *Handlers) DeleteFriend(ctx context.Context, req *api.FriendRequest) (*api.Empty, error) {
	userID, err := uuid.Parse(req.UserID.Id)
	if err != nil {
		h.logger.Error("Unable to parse uuid", zap.Error(err))
		return nil, err
	}

	friendID, err := uuid.Parse(req.FriendID.Id)
	if err != nil {
		h.logger.Error("Unable to parse uuid", zap.Error(err))
		return nil, err
	}

	err = h.processor.DeleteFriend(userID, friendID)
	if err != nil {
		h.logger.Error("Error occurred in call to processor", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}
