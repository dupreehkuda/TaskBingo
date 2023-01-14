package handlers

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errs "github.com/dupreehkuda/TaskBingo/task-data-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/task-data-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/task-data-service/pkg/api"
)

// AddOneTaskPack handles new task pack addition
func (h *Handlers) AddOneTaskPack(ctx context.Context, req *api.NewTaskPackRequest) (*api.NewTaskPackResponse, error) {
	data := models.TaskPack{
		TaskID: req.Id,
		Tasks:  req.Tasks,
	}

	err := h.processor.AddTaskPack(&data)

	switch {
	case err == errs.ErrPackAlreadyExists:
		return nil, status.Error(codes.AlreadyExists, "Exists")
	case err != nil:
		h.logger.Error("Error in call to processor", zap.Error(err))
		return nil, err
	}

	return &api.NewTaskPackResponse{Id: data.TaskID}, nil
}
