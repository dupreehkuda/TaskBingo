package processors

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (p processor) AcceptFriend(userID, friendID uuid.UUID) error {
	err := p.storage.AcceptFriend(userID, friendID)
	if err != nil {
		p.logger.Error("Error occurred in call to storage", zap.Error(err))
		return err
	}

	return nil
}

func (p processor) DeleteFriend(userID, friendID uuid.UUID) error {
	err := p.storage.DeleteFriend(userID, friendID)
	if err != nil {
		p.logger.Error("Error occurred in call to storage", zap.Error(err))
		return err
	}

	return nil
}

func (p processor) RequestFriend(userID, friendID uuid.UUID) error {
	err := p.storage.RequestFriend(userID, friendID)
	if err != nil {
		p.logger.Error("Error occurred in call to storage", zap.Error(err))
		return err
	}

	return nil
}
