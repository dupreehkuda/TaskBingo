package processors

import "go.uber.org/zap"

// RateTaskPack adds to pack rating
func (p processor) RateTaskPack(pack string) error {
	err := p.userStorage.RateTaskPack(pack)
	if err != nil {
		p.logger.Error("Error occurred calling user storage", zap.Error(err))
		return err
	}

	return nil
}

// UnrateTaskPack removes from pack rating
func (p processor) UnrateTaskPack(pack string) error {
	err := p.userStorage.UnrateTaskPack(pack)
	if err != nil {
		p.logger.Error("Error occurred calling user storage", zap.Error(err))
		return err
	}

	return nil
}
