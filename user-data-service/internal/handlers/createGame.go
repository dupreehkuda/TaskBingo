package handlers

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

// CreateGame handles the operation of writing new game to db
func (h *Handlers) CreateGame(ctx context.Context, req *api.NewGameRequest) (*api.Empty, error) {
	err := h.processor.CreateGame(&models.Game{
		GameID:       uuid.MustParse(req.GameID.Id),
		User1Id:      req.User1Id.Id,
		User2Id:      req.User2Id.Id,
		PackId:       req.Pack,
		Status:       req.Status,
		User1Bingo:   req.User1Bingo,
		User2Bingo:   req.User2Bingo,
		Winner:       req.Winner.Id,
		Numbers:      req.Numbers,
		User1Numbers: req.User1Numbers,
		User2Numbers: req.User2Numbers,
	})

	if err != nil {
		h.logger.Error("Error occurred calling processor", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, nil
}

// TODO implement AcceptGame
func (h *Handlers) AcceptGame(ctx context.Context, req *api.UUID) (*api.Empty, error) {
	return nil, nil
}
