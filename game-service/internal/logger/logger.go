package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

// InitializeLogger initializes new Logger instance
func InitializeLogger(localhost bool) *zap.Logger {
	if localhost {
		Logger, _ = zap.NewDevelopment()
	} else {
		Logger, _ = zap.NewProduction()
	}

	return Logger
}
