package taskRepository

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errs "github.com/dupreehkuda/TaskBingo/game-service/internal/customErrors"
	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// GetMultiplePacks retrieves multiple task packs from task service
func (t taskRepository) GetMultiplePacks(packIDs []string) (*[]models.TaskPack, error) {
	resp, err := t.conn.GetMultiplePacks(context.Background(), &api.GetMultiplePacksRequest{Ids: packIDs})

	statusCode, _ := status.FromError(err)
	if statusCode.Code() == codes.NotFound {
		return nil, errs.ErrNoSuchPack
	}

	var packs []models.TaskPack

	for _, val := range resp.Packs {
		packs = append(packs, models.TaskPack{
			ID: val.Id,
			Pack: models.Pack{
				Title: val.Title,
				Tasks: val.Tasks,
			},
		})
	}

	return &packs, nil
}
