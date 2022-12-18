package taskClient

import (
	"context"
	"encoding/json"

	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

func (t taskClient) SetTaskPack(pack *models.TaskPack) error {
	tasks, err := json.Marshal(pack)

	msg := &api.NewTaskPackRequest{}
	err = protojson.Unmarshal(tasks, msg)
	if err != nil {
		t.logger.Error("Error marshalling", zap.Error(err))
		return err
	}

	_, err = t.conn.AddOneTaskPack(context.Background(), msg)

	if err != nil {
		t.logger.Error("Error occurred in call to task service", zap.Error(err))
		return err
	}

	return nil
}
