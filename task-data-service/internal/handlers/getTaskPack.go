package handlers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errs "github.com/dupreehkuda/TaskBingo/task-data-service/internal/customErrors"
	api "github.com/dupreehkuda/TaskBingo/task-data-service/pkg/api"
)

// GetOneTaskPack handles getting one task pack operation
func (h *Handlers) GetOneTaskPack(ctx context.Context, req *api.TaskPackRequest) (*api.TaskPackResponse, error) {
	resp, err := h.processor.GetTaskPack(req.Id)

	switch {
	case err == errs.ErrNoSuchPack:
		return &api.TaskPackResponse{}, status.Error(codes.NotFound, "NSP")
	case err != nil:
		return &api.TaskPackResponse{}, err
	}

	return &api.TaskPackResponse{
		Id:    resp.TaskID,
		Tasks: resp.Tasks,
	}, nil
}
