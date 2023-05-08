package userRepository

import (
	"context"

	"go.uber.org/zap"
)

// AssignNewPack calls user service to assign new pack to creator and add it to ratings list
func (u userRepository) AssignNewPack(ctx context.Context, userID, packID, packName string) error {
	_, err := u.conn.AssignNewPack(ctx, mapToAssignNewPackRequest(userID, packID, packName))

	if err != nil {
		u.logger.Error("Error when assigning new pack", zap.Error(err))
		return err
	}

	return nil
}
