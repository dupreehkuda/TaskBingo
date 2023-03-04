package processors

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// AssignNewPack assigns fresh pack to creator and lists the pack in rating
func (p processor) AssignNewPack(userID, packID uuid.UUID, packName string) error {
	err := p.storage.AssignNewPack(userID, packID, packName)
	if err != nil {
		p.logger.Error("Error while retrieving data", zap.Error(err))
		return err
	}

	return nil
}
