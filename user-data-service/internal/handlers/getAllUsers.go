package handlers

import (
	"context"

	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

// GetAllPeople handles the operation of getting all people on service for the leaderboard
func (h *Handlers) GetAllPeople(ctx context.Context, _ *api.Empty) (*api.People, error) {
	users, err := h.service.GetAllUsers(ctx)
	if err != nil {
		h.logger.Error("Unable to call service", zap.Error(err))
		return nil, err
	}

	return mapToPeople(users), nil
}
