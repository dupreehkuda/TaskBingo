package processors

import "go.uber.org/zap"

// LikeTaskPack likes pack by user
func (p processor) LikeTaskPack(userID, pack string) error {
	err := p.userStorage.LikeTaskPack(userID, pack)
	if err != nil {
		p.logger.Error("Error occurred calling user storage", zap.Error(err))
		return err
	}

	return nil
}

// DislikeTaskPack dislikes pack by user
func (p processor) DislikeTaskPack(userID, pack string) error {
	err := p.userStorage.DislikeTaskPack(userID, pack)
	if err != nil {
		p.logger.Error("Error occurred calling user storage", zap.Error(err))
		return err
	}

	return nil
}
