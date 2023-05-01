package repository

import (
	"go.uber.org/zap"
)

// CheckPackExistence checks if task pack already exists
func (r repository) CheckPackExistence(taskId string) (bool, error) {
	_, err := r.handle.JSONGet(taskId, ".")
	if err != nil {
		if err.Error() == "redis: nil" {
			return false, nil
		}
		r.logger.Error("Error occurred when calling redis", zap.Error(err))
		return false, err
	}

	return true, nil
}
