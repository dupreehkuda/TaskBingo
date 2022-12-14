package processors

import "go.uber.org/zap"

// LikePack likes or dislikes the pack by inc
func (p processor) LikePack(login, pack string, inc int) error {
	err := p.storage.LikePack(login, pack, inc)
	if err != nil {
		p.logger.Error("Error in call to storage", zap.Error(err))
		return err
	}

	return nil
}
