package handlers

import (
	"context"

	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/user-data-service/pkg/api"
)

// GetRatedPacks handles the operation of getting some packs in desc rating order
func (h *Handlers) GetRatedPacks(ctx context.Context, _ *api.Empty) (*api.RatedPacksResponse, error) {
	resp, err := h.processor.GetRatedPacks()
	if err != nil {
		h.logger.Error("Unable to call processors", zap.Error(err))
		return nil, err
	}

	return &api.RatedPacksResponse{Packs: resp}, nil
}
