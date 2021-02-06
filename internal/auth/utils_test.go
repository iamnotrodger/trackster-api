package auth

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
)

func TestGenerateAndVerify(t *testing.T) {
	tokenString, err := generateToken("1")
	if err != nil {
		t.Error(err)
	}

	t.Log(tokenString)

	claims, err := verifyToken(tokenString)
	if err != nil {
		t.Error(err)
	}

	userID := claims.(jwt.MapClaims)["user_id"].(string)

	t.Log(userID)

}
