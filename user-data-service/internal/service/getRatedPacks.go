package service

import (
	"context"

	"go.uber.org/zap"
)

// GetRatedPacks gets some of most rated packs
func (s service) GetRatedPacks(ctx context.Context) ([]string, error) {
	packs, err := s.repository.GetRatedPacks(ctx)
	if err != nil {
		s.logger.Error("Error in call to repository", zap.Error(err))
		return nil, err
	}

	return packs, nil
}
