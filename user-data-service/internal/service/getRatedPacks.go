package service

import (
	"go.uber.org/zap"
)

// GetRatedPacks gets some of most rated packs
func (s service) GetRatedPacks() ([]string, error) {
	packs, err := s.repository.GetRatedPacks()
	if err != nil {
		s.logger.Error("Error in call to repository", zap.Error(err))
		return nil, err
	}

	return packs, nil
}
