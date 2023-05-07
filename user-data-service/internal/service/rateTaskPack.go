package service

import (
	"context"

	"go.uber.org/zap"
)

// RatePack rates pack by inc
func (s service) RatePack(ctx context.Context, userID, pack string, inc int) error {
	err := s.repository.RatePack(ctx, userID, pack, inc)
	if err != nil {
		s.logger.Error("Error in call to repository", zap.Error(err))
		return err
	}

	return nil
}
