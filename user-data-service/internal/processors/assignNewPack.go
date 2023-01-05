package processors

import "go.uber.org/zap"

// AssignNewPack assigns fresh pack to creator and lists the pack in rating
func (p processor) AssignNewPack(login, pack string) error {
	err := p.storage.AssignNewPack(login, pack)
	if err != nil {
		p.logger.Error("Error while retrieving data", zap.Error(err))
		return err
	}

	return nil
}
