package middleware

import (
	"net/http"

	"go.uber.org/zap"
)

// Middleware is an interface for middleware layer
type Middleware interface {
	CheckToken(next http.Handler) http.Handler
}

// middleware provides services middleware
type middleware struct {
	logger *zap.Logger
}

// New creates new instance of middleware
func New(logger *zap.Logger) *middleware {
	return &middleware{
		logger: logger,
	}
}
