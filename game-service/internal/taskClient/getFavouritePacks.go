package taskClient

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// GetMultiplePacks retrieves multiple task packs from task service
func (t taskClient) GetMultiplePacks(packIDs []string) (*[]models.TaskPack, error) {
	resp, err := t.conn.GetMultiplePacks(context.Background(), &api.GetMultiplePacksRequest{Id: packIDs})

	statusCode, _ := status.FromError(err)
	if statusCode.Code() == codes.NotFound {
		return nil, errs.ErrNoSuchPack
	}

	var packs []models.TaskPack

	for _, val := range resp.Packs {
		packs = append(packs, models.TaskPack{
			TaskID: val.Id,
			Tasks:  val.Tasks,
		})
	}

	return &packs, nil
}
