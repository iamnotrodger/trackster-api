package auth

import (
	"os"
	"testing"
)

func TestGenerateAndVerify(t *testing.T) {
	t.Log(os.Getenv("ACCESS_TOKEN_SECRET"))
	var signingKey = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	t.Log(signingKey)
}
