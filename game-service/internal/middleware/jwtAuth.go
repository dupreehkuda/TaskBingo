package middleware

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/game-service/internal/models"
)

// CheckToken implements JWT token parsing and authorizing users
func (m middleware) CheckToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("auth")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			m.logger.Error("Error occurred while unpacking cookies", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var signingKey = []byte(os.Getenv("secret"))

		if token != nil {
			token, err := jwt.Parse(token.Value, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					w.WriteHeader(http.StatusUnauthorized)
					_, err := w.Write([]byte("You're Unauthorized"))
					if err != nil {
						return "", err
					}
					return "", err
				}

				return signingKey, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				_, err := w.Write([]byte("You're Unauthorized due to error"))
				if err != nil {
					m.logger.Error("Unable to write response", zap.Error(err))
				}
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				var ctxKey models.LoginKey = "login"
				login := claims["user"].(string)
				ctx := context.WithValue(r.Context(), ctxKey, login)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}

		w.WriteHeader(http.StatusUnauthorized)
		_, err = w.Write([]byte("You're Unauthorized"))
		if err != nil {
			m.logger.Error("Unable to write response", zap.Error(err))
			return
		}
	})
}
