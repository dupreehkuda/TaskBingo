package handlers

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

// AssignNewPack handles the operation of assigning fresh pack to rating list and tie it to the creator
func (h *Handlers) AssignNewPack(ctx context.Context, req *api.AssignNewPackRequest) (*api.Empty, error) {
	userID, err := uuid.Parse(req.UserID.Id)
	if err != nil {
		h.logger.Error("Unable to parse uuid", zap.Error(err))
		return nil, err
	}

	packID, err := uuid.Parse(req.PackID.Id)
	if err != nil {
		h.logger.Error("Unable to parse uuid", zap.Error(err))
		return nil, err
	}

	err = h.processor.AssignNewPack(userID, packID, req.PackName)
	if err != nil {
		h.logger.Error("Error occurred in call to processor", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, err
}
