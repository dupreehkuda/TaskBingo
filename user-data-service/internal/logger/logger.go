package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

// InitializeLogger initializes new Logger instance
func InitializeLogger() *zap.Logger {
	Logger, _ = zap.NewDevelopment()
	return Logger
}
