package processors

import (
	"go.uber.org/zap"
)

func (p processor) RequestFriend(login, person string) error {
	err := p.userStorage.RequestFriend(login, person)
	if err != nil {
		p.logger.Error("Error occurred in call to user service", zap.Error(err))
		return err
	}

	return nil
}

func (p processor) AcceptFriend(login, person string) error {
	err := p.userStorage.AcceptFriend(login, person)
	if err != nil {
		p.logger.Error("Error occurred in call to user service", zap.Error(err))
		return err
	}

	return nil
}

func (p processor) DeleteFriend(login, person string) error {
	err := p.userStorage.DeleteFriend(login, person)
	if err != nil {
		p.logger.Error("Error occurred in call to user service", zap.Error(err))
		return err
	}

	return nil
}
