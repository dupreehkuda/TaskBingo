package taskClient

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// GetTaskPack retrieves a task pack from task service
func (t taskClient) GetTaskPack(packID string) (*models.TaskPack, error) {
	resp, err := t.conn.GetOneTaskPack(context.Background(), &api.TaskPackRequest{Id: packID})

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
