package processors

import (
	"go.uber.org/zap"
)

// GetRatedPacks gets some of most rated packs
func (p processor) GetRatedPacks() ([]string, error) {
	packs, err := p.storage.GetRatedPacks()
	if err != nil {
		p.logger.Error("Error in call to storage", zap.Error(err))
		return nil, err
	}

	return packs, nil
}
