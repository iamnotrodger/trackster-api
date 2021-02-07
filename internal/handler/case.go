package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/iamnotrodger/trackster-api/internal/model"
	"github.com/jmoiron/sqlx"
)

//CasePost struct
type CasePost struct {
	CaseTime string `json:"case_time"`
}

// NotifyUser that they were exposed to the rona
func NotifyUser(user model.User) error {
	// TODO: send them and email
	return errors.New("Not implemented")
}

//PostCase func
func PostCase(db *sqlx.DB) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var casePost CasePost
		err := json.NewDecoder(r.Body).Decode(&casePost)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		tos, err := parseStringToTime(casePost.CaseTime)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		userID := r.Header.Get("user_id")
		caseStruct := model.Case{UserID: userID, TimeOfSymptoms: tos}
		err = caseStruct.Insert(db)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		timeUntilContagious := tos.AddDate(0, 0, 14).Format("2006-01-02T15:04:05.000Z")
		contacts, err := model.SelectContactByTmeInterval(db, timeUntilContagious, casePost.CaseTime)
		if err != nil && len(contacts) != 0 {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		validTraces := func(userID string, contacts []model.Contact) []model.Contact {
			riskyPlaces := make(map[string]bool)
			for _, v := range contacts {
				if userID == v.UserID {
					riskyPlaces[v.LocationID] = true
				}
			}

			vc := make([]model.Contact, 0)
			for _, v := range contacts {
				res, ok := riskyPlaces[v.LocationID]
				if v.UserID != userID && res && ok {
					vc = append(vc, v)
				}
			}
			return vc
		}(userID, contacts)

		r.Header.Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(validTraces)

	})
}
