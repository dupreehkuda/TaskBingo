package userClient

import (
	"context"

	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// AssignNewPack calls user service to assign new pack to creator and add it to ratings list
func (u userClient) AssignNewPack(userID, pack string) error {
	_, err := u.conn.AssignNewPack(context.Background(), &api.AssignNewPackRequest{
		UserID: &api.UUID{Id: userID},
		Pack:   pack,
	})

	if err != nil {
		u.logger.Error("Error when assigning new pack", zap.Error(err))
		return err
	}

	return nil
}
