package processors

import "go.uber.org/zap"

// RatePack rates pack by inc
func (p processor) RatePack(login, pack string, inc int) error {
	err := p.storage.RatePack(login, pack, inc)
	if err != nil {
		p.logger.Error("Error in call to storage", zap.Error(err))
		return err
	}

	return nil
}
