package tokens_test

import (
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"

	"github.com/dupreehkuda/TaskBingo/pkg/tokens"
)

func TestGenerateJWT_RoundTrip(t *testing.T) {
	secret := "test-secret"
	raw, err := tokens.GenerateJWT("u-1", "alice", secret)
	require.NoError(t, err)
	require.NotEmpty(t, raw)

	parsed, err := jwt.Parse(raw, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	require.NoError(t, err)

	claims := parsed.Claims.(jwt.MapClaims)
	require.Equal(t, "u-1", claims["userID"])
	require.Equal(t, "alice", claims["username"])
	require.NotZero(t, claims["exp"])
}
