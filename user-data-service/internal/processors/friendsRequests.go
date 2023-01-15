package processors

import "go.uber.org/zap"

func (p processor) AcceptFriend(login, person string) error {
	err := p.storage.AcceptFriend(login, person)
	if err != nil {
		p.logger.Error("Error occurred in call to storage", zap.Error(err))
		return err
	}

	return nil
}

func (p processor) DeleteFriend(login, person string) error {
	err := p.storage.DeleteFriend(login, person)
	if err != nil {
		p.logger.Error("Error occurred in call to storage", zap.Error(err))
		return err
	}

	return nil
}

func (p processor) RequestFriend(login, person string) error {
	err := p.storage.RequestFriend(login, person)
	if err != nil {
		p.logger.Error("Error occurred in call to storage", zap.Error(err))
		return err
	}

	return nil
}
