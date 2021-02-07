package handler

import (
	"encoding/json"
	"net/http"

	"github.com/iamnotrodger/trackster-api/internal/model"
	"github.com/jmoiron/sqlx"
)

type UserPost struct {
	Email       string `json:"email"`
}

func PostUser(db *sqlx.DB) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		var payload UserPost
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			newUser := model.User{Email: payload.Email}
			err := newUser.Insert(db)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(newUser)
			}
		}
	})
}
