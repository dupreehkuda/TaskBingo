package handlers

import (
	"context"

	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

func (h *Handlers) AchieveGame(ctx context.Context, req *api.GameRequest) (*api.Empty, error) {
	err := h.service.AchieveGame(ctx, &models.Game{
		GameID:       req.GameID,
		User1Id:      req.User1Id,
		User2Id:      req.User2Id,
		PackId:       req.Pack,
		Status:       req.Status,
		User1Bingo:   req.User1Bingo,
		User2Bingo:   req.User2Bingo,
		Winner:       req.Winner,
		Numbers:      req.Numbers,
		User1Numbers: req.User1Numbers,
		User2Numbers: req.User2Numbers,
	})

	if err != nil {
		h.logger.Error("Error occurred calling service", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}

func (h *Handlers) GetGame(ctx context.Context, req *api.GetGameRequest) (*api.GameRequest, error) {
	game, err := h.service.GetGame(ctx, req.GameID)
	if err != nil {
		h.logger.Error("Error occurred calling service", zap.Error(err))
		return nil, err
	}

	return &api.GameRequest{
		GameID:       game.GameID,
		User1Id:      game.User1Id,
		User2Id:      game.User2Id,
		Pack:         game.PackId,
		Status:       game.Status,
		User1Bingo:   game.User1Bingo,
		User2Bingo:   game.User2Bingo,
		Winner:       game.Winner,
		Numbers:      game.Numbers,
		User1Numbers: game.User1Numbers,
		User2Numbers: game.User2Numbers,
	}, nil
}
