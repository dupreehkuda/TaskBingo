package middleware

import (
	"net/http"

	"go.uber.org/zap"
)

type Middleware interface {
	CheckToken(next http.Handler) http.Handler
	AllowUpgradeHeaders(next http.Handler) http.Handler
	CheckCompression(next http.Handler) http.Handler
	WriteCompressed(next http.Handler) http.Handler
	RequestLogger(next http.Handler) http.Handler
}

type middleware struct {
	logger    *zap.Logger
	jwtSecret string
}

func New(logger *zap.Logger, jwtSecret string) Middleware {
	return &middleware{logger: logger, jwtSecret: jwtSecret}
}
