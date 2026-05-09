package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/internal/models"
)

func (m middleware) CheckToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			m.logger.Error("Error unpacking cookie", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		token, err := jwt.Parse(cookie.Value, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(m.jwtSecret), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("You're Unauthorized due to error"))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("You're Unauthorized"))
			return
		}

		userID, _ := claims["userID"].(string)
		username, _ := claims["username"].(string)
		if userID == "" {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("You're Unauthorized"))
			return
		}

		ctx := context.WithValue(r.Context(), models.CtxUserIDKey, userID)
		ctx = context.WithValue(ctx, models.CtxUsernameKey, username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
