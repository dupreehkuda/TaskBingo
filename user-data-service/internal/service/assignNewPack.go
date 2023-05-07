package service

import (
	"context"

	"go.uber.org/zap"
)

// AssignNewPack assigns fresh pack to creator and lists the pack in rating
func (s service) AssignNewPack(ctx context.Context, userID, packID string, packName string) error {
	err := s.repository.AssignNewPack(ctx, userID, packID, packName)
	if err != nil {
		s.logger.Error("Error while retrieving data", zap.Error(err))
		return err
	}

	return nil
}
