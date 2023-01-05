package userClient

import (
	"context"

	"go.uber.org/zap"

	api "github.com/dupreehkuda/TaskBingo/game-service/pkg/api"
)

// GetRatedPacks retrieves most rated packs from user service
func (u userClient) GetRatedPacks() ([]string, error) {
	resp, err := u.conn.GetRatedPacks(context.Background(), &api.Empty{})
	if err != nil {
		u.logger.Error("Error when getting user data", zap.Error(err))
		return nil, err
	}

	return resp.Packs, nil
}
