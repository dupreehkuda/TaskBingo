package storage

import (
	"go.uber.org/zap"
)

// CheckPackExistence checks if task pack already exists
func (s storage) CheckPackExistence(taskId string) (bool, error) {
	_, err := s.handle.JSONGet(taskId, ".")
	if err != nil {
		if err.Error() == "redis: nil" {
			return false, nil
		}
		s.logger.Error("Error occurred when calling redis", zap.Error(err))
		return false, err
	}

	return true, nil
}
