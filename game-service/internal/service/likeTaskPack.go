package service

import "go.uber.org/zap"

// LikeTaskPack likes pack by user
func (s service) LikeTaskPack(userID, pack string) error {
	err := s.userRepository.LikeTaskPack(userID, pack)
	if err != nil {
		s.logger.Error("Error occurred calling user repository", zap.Error(err))
		return err
	}

	return nil
}

// DislikeTaskPack dislikes pack by user
func (s service) DislikeTaskPack(userID, pack string) error {
	err := s.userRepository.DislikeTaskPack(userID, pack)
	if err != nil {
		s.logger.Error("Error occurred calling user repository", zap.Error(err))
		return err
	}

	return nil
}
