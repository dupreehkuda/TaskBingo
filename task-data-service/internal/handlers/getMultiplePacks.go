package handlers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errs "github.com/dupreehkuda/TaskBingo/task-data-service/internal/customErrors"
	api "github.com/dupreehkuda/TaskBingo/task-data-service/pkg/api"
)

// GetMultiplePacks retrieves multiple task packs from the database
func (h *Handlers) GetMultiplePacks(ctx context.Context, req *api.GetMultiplePacksRequest) (*api.GetMultiplePacksResponse, error) {
	resp, err := h.processor.GetMultiplePacks(req.Ids)

	switch {
	case err == errs.ErrNoSuchPack:
		return &api.GetMultiplePacksResponse{}, status.Error(codes.NotFound, "NSP")
	case err != nil:
		return &api.GetMultiplePacksResponse{}, err
	}

	var ans []*api.TaskPackResponse
	for _, val := range *resp {
		ans = append(ans, &api.TaskPackResponse{
			Id:    val.ID,
			Title: val.Pack.Title,
			Tasks: val.Pack.Tasks,
		})
	}

	return &api.GetMultiplePacksResponse{Packs: ans}, nil
}
