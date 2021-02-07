package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/iamnotrodger/trackster-api/internal/auth"
)

//RefreshToken assigns a new access-token using refresh-token
func RefreshToken(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	refreshToken := cookie.Value
	fmt.Println(refreshToken)

	claims, err := auth.VerifyRefreshToken(refreshToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	userID := claims.(jwt.MapClaims)["user_id"].(string)
	fmt.Println(userID)

	accessToken, err := auth.GenerateAccessToken(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		AccessToken string `json:"access_token"`
	}{accessToken})
}
