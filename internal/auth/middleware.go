package auth

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

//Middleware verifies that the jwt exist and is not expired
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := verifyAccessToken(tokenString)
		if err != nil {
			http.Error(w, "Token Unauthorized", http.StatusUnauthorized)
			return
		}

		userID := claims.(jwt.MapClaims)["user_id"].(string)
		r.Header.Set(("user_id"), userID)

		next.ServeHTTP(w, r)
	})
}
