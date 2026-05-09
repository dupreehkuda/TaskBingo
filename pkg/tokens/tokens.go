package tokens

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// TTL is how long a freshly minted token stays valid.
const TTL = 72 * time.Hour

// GenerateJWT signs a token with the given secret and embeds userID/username claims.
// Note: the env var was previously read inline as "secret" (lowercase). The caller
// is now responsible for providing the value, which makes the helper testable and
// keeps Config the single source of truth.
func GenerateJWT(userID, username, secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(TTL).Unix()
	claims["userID"] = userID
	claims["username"] = username
	return token.SignedString([]byte(secret))
}
