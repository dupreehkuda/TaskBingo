package service

import (
	"go.uber.org/zap"
)

// LikePack likes or dislikes the pack by inc
func (s service) LikePack(userID, pack string, inc int) error {
	err := s.repository.LikePack(userID, pack, inc)
	if err != nil {
		s.logger.Error("Error in call to repository", zap.Error(err))
		return err
	}

	return nil
}
