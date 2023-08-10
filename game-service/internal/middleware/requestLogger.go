package middleware

import (
	"net/http"

	"go.uber.org/zap"
)

func (m middleware) RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.logger.Debug("Request",
			zap.Any("method", r.Method),
			zap.Any("host", r.Host),
			zap.Any("remote addr", r.RemoteAddr),
			zap.Any("req uri", r.RequestURI),
			zap.Any("header", r.Header),
		)

		next.ServeHTTP(w, r)
	})
}
