package taskRepository

import (
	"context"
	"encoding/json"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// SetTaskPack sets a new task pack in task service
func (t taskRepository) SetTaskPack(pack *models.TaskPack) error {
	tasks, err := json.Marshal(pack)

	msg := &api.NewTaskPackRequest{}
	err = protojson.Unmarshal(tasks, msg)
	if err != nil {
		t.logger.Error("Error marshalling", zap.Error(err))
		return err
	}

	_, err = t.conn.AddOneTaskPack(context.Background(), msg)

	statusCode, _ := status.FromError(err)
	if statusCode.Code() == codes.NotFound {
		return errs.ErrPackAlreadyExists
	}

	if err != nil {
		t.logger.Error("Error occurred in call to task service", zap.Error(err))
		return err
	}

	return nil
}
