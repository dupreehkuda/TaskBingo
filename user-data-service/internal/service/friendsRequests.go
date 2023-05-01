package service

import (
	"go.uber.org/zap"
)

func (s service) AcceptFriend(userID, friendID string) error {
	err := s.repository.AcceptFriend(userID, friendID)
	if err != nil {
		s.logger.Error("Error occurred in call to repository", zap.Error(err))
		return err
	}

	return nil
}

func (s service) DeleteFriend(userID, friendID string) error {
	err := s.repository.DeleteFriend(userID, friendID)
	if err != nil {
		s.logger.Error("Error occurred in call to repository", zap.Error(err))
		return err
	}

	return nil
}

func (s service) RequestFriend(userID, friendID string) error {
	err := s.repository.RequestFriend(userID, friendID)
	if err != nil {
		s.logger.Error("Error occurred in call to repository", zap.Error(err))
		return err
	}

	return nil
}
