package handler

//TODO: login
//TODO: refresh token

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/iamnotrodger/trackster-api/internal/auth"
	"github.com/iamnotrodger/trackster-api/internal/model"
	"github.com/jmoiron/sqlx"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/api/auth/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)

//Login for testing
func Login(w http.ResponseWriter, r *http.Request) {
	var htmlIndex = `
		<html>
			<body>
			<a href="/api/auth/google">Google Log In</a>
			</body>
		</html>`

	fmt.Fprintf(w, htmlIndex)
}

//GoogleLogin redirects to google's login
func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

//GoogleCallback func
func GoogleCallback(db *sqlx.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		var userID string

		//Checks if the user if the user is registered else, if not the user is registered
		registeredUser, err := model.SelectUserByProviderID(db, user.ProviderID)
		if err != nil || (model.User{}) == *registeredUser {
			err := user.Insert(db)
			if err != nil {
				http.Error(w, "Unable to Register User: \n"+err.Error(), http.StatusInternalServerError)
				return
			}
			userID = user.UserID
		} else {
			userID = registeredUser.UserID
		}

		accessToken, err := auth.GenerateAccessToken(userID)
		if err != nil {
			http.Error(w, "Failed to Generate JWT: \n"+err.Error(), http.StatusInternalServerError)
		}
		refreshToken, err := auth.GenerateRefreshToken(userID)
		if err != nil {
			http.Error(w, "Failed to Generate JWT: \n"+err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")

		cookie := http.Cookie{
			Name:     "refresh_token",
			Value:    refreshToken,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, &cookie)

		json.NewEncoder(w).Encode(struct {
			AccessToken string `json:"access_token"`
		}{accessToken})

	})
}

func getUserInfo(state string, code string) (*model.User, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	var user model.User
	err = json.NewDecoder(response.Body).Decode(&user)

	return &user, nil
}
