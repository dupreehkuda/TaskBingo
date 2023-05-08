package handlers

import (
	"context"

	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

// AchieveGame handles the operation of writing finished game to repository
func (h *Handlers) AchieveGame(ctx context.Context, req *api.GameRequest) (*api.Empty, error) {
	err := h.service.AchieveGame(ctx, mapFromGameRequest(req))

	if err != nil {
		h.logger.Error("Error occurred calling service", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}

// GetGame handles the operation of getting current game
func (h *Handlers) GetGame(ctx context.Context, req *api.GetGameRequest) (*api.GameRequest, error) {
	game, err := h.service.GetGame(ctx, req.GameID)
	if err != nil {
		h.logger.Error("Error occurred calling service", zap.Error(err))
		return nil, err
	}

	return mapToGameRequest(game), nil
}
