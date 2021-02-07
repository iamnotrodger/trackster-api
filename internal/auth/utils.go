package auth

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//GenerateAccessToken generate a jwt token containing user ID
func GenerateAccessToken(userID string) (string, error) {
	var signingKey = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

//GenerateRefreshToken generate a jwt token containing user ID
func GenerateRefreshToken(userID string) (string, error) {
	var signingKey = []byte(os.Getenv("REFRESH_TOKEN_SECRET"))
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().AddDate(1, 0, 0).Unix()

	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func verifyAccessToken(tokenString string) (jwt.Claims, error) {
	var signingKey = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	return token.Claims, err
}

//VerifyRefreshToken verifies the refresh-token cookie
func VerifyRefreshToken(tokenString string) (jwt.Claims, error) {
	var signingKey = []byte(os.Getenv("REFRESH_TOKEN_SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	return token.Claims, err
}
