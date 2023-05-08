package service

import (
	"go.uber.org/zap"
)

// RequestFriend calls user repository to request friendship
func (s service) RequestFriend(userID, friendID string) error {
	err := s.userRepository.RequestFriend(userID, friendID)
	if err != nil {
		s.logger.Error("Error occurred in call to user service", zap.Error(err))
		return err
	}

	return nil
}

// AcceptFriend calls user repository to accept friendship
func (s service) AcceptFriend(userID, friendID string) error {
	err := s.userRepository.AcceptFriend(userID, friendID)
	if err != nil {
		s.logger.Error("Error occurred in call to user service", zap.Error(err))
		return err
	}

	return nil
}

// DeleteFriend calls user repository to delete/cancel friendship
func (s service) DeleteFriend(userID, friendID string) error {
	err := s.userRepository.DeleteFriend(userID, friendID)
	if err != nil {
		s.logger.Error("Error occurred in call to user service", zap.Error(err))
		return err
	}

	return nil
}
