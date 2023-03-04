package tokens

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// GenerateJWT returns a JWT token
func GenerateJWT(userID, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(72 * time.Hour).Unix()
	claims["userID"] = userID
	claims["username"] = username
	tokenString, err := token.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		log.Print(err)
		return "Signing Error", err
	}

	return tokenString, nil
}
