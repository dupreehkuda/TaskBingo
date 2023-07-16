package middleware

import "net/http"

func (m middleware) AllowUpgradeHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Add("Connection", "upgrade")
		r.Header.Add("Upgrade", "websocket")

		next.ServeHTTP(w, r)
	})
}
