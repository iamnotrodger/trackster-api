package auth

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//GenerateToken generate a jwt token containing user ID
func GenerateToken(userID string) (string, error) {
	var signingKey = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	var signingKey = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	return token.Claims, err
}
