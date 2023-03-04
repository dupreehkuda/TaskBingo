package processors

import (
	"go.uber.org/zap"
)

func (p processor) RequestFriend(userID, friendID string) error {
	err := p.userStorage.RequestFriend(userID, friendID)
	if err != nil {
		p.logger.Error("Error occurred in call to user service", zap.Error(err))
		return err
	}

	return nil
}

func (p processor) AcceptFriend(userID, friendID string) error {
	err := p.userStorage.AcceptFriend(userID, friendID)
	if err != nil {
		p.logger.Error("Error occurred in call to user service", zap.Error(err))
		return err
	}

	return nil
}

func (p processor) DeleteFriend(userID, friendID string) error {
	err := p.userStorage.DeleteFriend(userID, friendID)
	if err != nil {
		p.logger.Error("Error occurred in call to user service", zap.Error(err))
		return err
	}

	return nil
}
