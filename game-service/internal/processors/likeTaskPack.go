package processors

import "go.uber.org/zap"

// LikeTaskPack likes pack by user
func (p processor) LikeTaskPack(login, pack string) error {
	err := p.userStorage.LikeTaskPack(login, pack)
	if err != nil {
		p.logger.Error("Error occurred calling user storage", zap.Error(err))
		return err
	}

	return nil
}

// DislikeTaskPack dislikes pack by user
func (p processor) DislikeTaskPack(login, pack string) error {
	err := p.userStorage.DislikeTaskPack(login, pack)
	if err != nil {
		p.logger.Error("Error occurred calling user storage", zap.Error(err))
		return err
	}

	return nil
}
