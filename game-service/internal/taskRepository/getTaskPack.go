package taskRepository

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// GetTaskPack retrieves a task pack from task service
func (t taskRepository) GetTaskPack(ctx context.Context, packID string) (*models.TaskPack, error) {
	resp, err := t.conn.GetOneTaskPack(ctx, &api.TaskPackRequest{Id: packID})

	statusCode, _ := status.FromError(err)
	if statusCode.Code() == codes.NotFound {
		return nil, errs.ErrNoSuchPack
	}

	tasks := models.TaskPack{
		ID: resp.Id,
		Pack: models.Pack{
			Title: resp.Title,
			Tasks: resp.Tasks,
		},
	}

	return &tasks, nil
}
