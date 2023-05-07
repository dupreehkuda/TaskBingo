package service

import (
	"context"

	"go.uber.org/zap"
)

// LikePack likes or dislikes the pack by inc
func (s service) LikePack(ctx context.Context, userID, pack string, inc int) error {
	err := s.repository.LikePack(ctx, userID, pack, inc)
	if err != nil {
		s.logger.Error("Error in call to repository", zap.Error(err))
		return err
	}

	return nil
}
