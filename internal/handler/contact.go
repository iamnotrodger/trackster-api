package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/iamnotrodger/trackster-api/internal/model"
	"github.com/jmoiron/sqlx"
)

//ContactPost struct
type ContactPost struct {
	Address     string    `json:"address"`
	ContactTime time.Time `json:"contact_time"`
}

//PostContact lol
func PostContact(db *sqlx.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var contactPost ContactPost
		err := json.NewDecoder(r.Body).Decode(&contactPost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		location, err := model.SelectLocationByAddress(db, contactPost.Address)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if (model.Location{}) == location {
			location.Address = contactPost.Address
			location.Insert(db)
		}

		var contact model.Contact
		contact.ContactTime = contactPost.ContactTime

		err = contact.Insert(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(contact)
	})
}
