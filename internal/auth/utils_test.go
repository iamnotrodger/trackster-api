package auth

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
)

func TestGenerateAndVerify(t *testing.T) {
	tokenString, err := GenerateAccessToken("1")
	if err != nil {
		t.Error(err)
	}

	t.Log(tokenString)

	claims, err := verifyAccessToken(tokenString)
	if err != nil {
		t.Error(err)
	}

	userID := claims.(jwt.MapClaims)["user_id"].(string)

	t.Log(userID)

}
