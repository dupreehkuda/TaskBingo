package middleware

import (
	"net/http"

	"go.uber.org/zap"
)

func (m middleware) RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.logger.Debug("Request",
			zap.String("method", r.Method),
			zap.String("host", r.Host),
			zap.String("remote_addr", r.RemoteAddr),
			zap.String("uri", r.RequestURI),
			zap.Any("header", r.Header),
		)
		next.ServeHTTP(w, r)
	})
}
