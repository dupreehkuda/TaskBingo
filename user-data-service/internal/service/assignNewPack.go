package service

import (
	"go.uber.org/zap"
)

// AssignNewPack assigns fresh pack to creator and lists the pack in rating
func (s service) AssignNewPack(userID, packID string, packName string) error {
	err := s.repository.AssignNewPack(userID, packID, packName)
	if err != nil {
		s.logger.Error("Error while retrieving data", zap.Error(err))
		return err
	}

	return nil
}
