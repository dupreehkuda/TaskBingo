package taskClient

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

func (t taskClient) GetTaskPack(packID string) (*models.TaskPack, error) {
	resp, err := t.conn.GetOneTaskPack(context.Background(), &api.TaskPackRequest{Id: packID})

	statusCode, _ := status.FromError(err)
	if statusCode.Code() == codes.NotFound {
		return nil, errs.ErrNoSuchPack
	}

	t.logger.Debug("code", zap.String("received code", statusCode.Code().String()))

	tasks := models.TaskPack{
		TaskID: resp.Id,
		Tasks:  resp.Tasks,
	}

	return &tasks, nil
}
