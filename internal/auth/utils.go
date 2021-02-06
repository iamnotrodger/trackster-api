package auth

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var signingKey = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))

// GenerateToken function
func GenerateToken(userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodES256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

// VerifyToken function
func VerifyToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
