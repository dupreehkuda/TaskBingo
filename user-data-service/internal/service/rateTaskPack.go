package service

import (
	"go.uber.org/zap"
)

// RatePack rates pack by inc
func (s service) RatePack(userID, pack string, inc int) error {
	err := s.repository.RatePack(userID, pack, inc)
	if err != nil {
		s.logger.Error("Error in call to repository", zap.Error(err))
		return err
	}

	return nil
}
