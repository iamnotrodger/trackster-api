package handler

import (
	"encoding/json"
	"net/http"

	"github.com/iamnotrodger/trackster-api/internal/model"
	"github.com/jmoiron/sqlx"
)

//ContactPost struct
type ContactPost struct {
	Address     string `json:"address"`
	ContactTime string `json:"contact_time"`
}

//PostContact lol
func PostContact(db *sqlx.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var contactPost ContactPost
		err := json.NewDecoder(r.Body).Decode(&contactPost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		location, err := model.SelectLocationByAddress(db, contactPost.Address)
		if err != nil || (model.Location{}) == location {
			location.Address = contactPost.Address

			err = location.Insert(db)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		t, err := parseStringToTime(contactPost.ContactTime)
		if err != nil {
			http.Error(w, "Bad Request: Invalid Time Format \n"+err.Error(), http.StatusBadRequest)
			return
		}

		var contact model.Contact
		//THIS IS JUST FOR TESTING FOR NOW
		contact.UserID = "1"
		contact.ContactTime = t
		contact.LocationID = location.LocationID

		err = contact.Insert(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(contact)
	})
}
