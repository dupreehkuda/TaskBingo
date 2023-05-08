package handlers

import (
	"context"

	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

// CreateGame handles the operation of writing new game to db
func (h *Handlers) CreateGame(ctx context.Context, req *api.GameRequest) (*api.Empty, error) {
	err := h.service.CreateGame(ctx, mapFromGameRequest(req))

	if err != nil {
		h.logger.Error("Error occurred calling service", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}

// AcceptGame handles the operation of accepting the game
func (h *Handlers) AcceptGame(ctx context.Context, req *api.StatusGameRequest) (*api.Empty, error) {
	if err := h.service.AcceptGame(ctx, req.UserID, req.GameID); err != nil {
		h.logger.Error("Error occurred calling service", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}

// DeleteGame handles the operation of deleting the game
func (h *Handlers) DeleteGame(ctx context.Context, req *api.StatusGameRequest) (*api.Empty, error) {
	if err := h.service.DeleteGame(ctx, req.UserID, req.GameID); err != nil {
		h.logger.Error("Error occurred calling service", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}
