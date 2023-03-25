package handlers

import (
	"context"

	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

// AssignNewPack handles the operation of assigning fresh pack to rating list and tie it to the creator
func (h *Handlers) AssignNewPack(ctx context.Context, req *api.AssignNewPackRequest) (*api.Empty, error) {
	err := h.processor.AssignNewPack(req.UserID, req.PackID, req.PackName)
	if err != nil {
		h.logger.Error("Error occurred in call to processor", zap.Error(err))
		return nil, err
	}

	return &api.Empty{}, err
}
