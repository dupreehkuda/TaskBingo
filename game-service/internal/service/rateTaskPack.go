package service

import "go.uber.org/zap"

// RateTaskPack adds to pack rating
func (s service) RateTaskPack(userID, pack string) error {
	err := s.userRepository.RateTaskPack(userID, pack)
	if err != nil {
		s.logger.Error("Error occurred calling user repository", zap.Error(err))
		return err
	}

	return nil
}

// UnrateTaskPack removes from pack rating
func (s service) UnrateTaskPack(userID, pack string) error {
	err := s.userRepository.UnrateTaskPack(userID, pack)
	if err != nil {
		s.logger.Error("Error occurred calling user repository", zap.Error(err))
		return err
	}

	return nil
}
