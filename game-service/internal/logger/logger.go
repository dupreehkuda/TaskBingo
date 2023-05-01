package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

// InitializeLogger initializes new Logger instance
func InitializeLogger(localhost bool) *zap.Logger {
	switch {
	case localhost:
		Logger, _ = zap.NewDevelopment()
	default:
		Logger, _ = zap.NewProduction()
	}

	return Logger
}
