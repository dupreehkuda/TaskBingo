package middleware

import "go.uber.org/zap"

type middleware struct {
	logger *zap.Logger
}

// New creates new instance of middleware
func New(logger *zap.Logger) *middleware {
	return &middleware{
		logger: logger,
	}
}
