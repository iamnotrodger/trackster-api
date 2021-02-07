package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/iamnotrodger/trackster-api/internal/model"
	"github.com/jmoiron/sqlx"
)

type CasePost struct {
	CaseTime time.Time `json:"case_time"`
}

// NotifyUser that they were exposed to the rona
func NotifyUser(user model.User) error {
	// TODO: send them and email
	return errors.New("Not implemented")
}

func PostCase(db *sqlx.DB) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// TODO: get user id from header
		var casePost CasePost
		err := json.NewDecoder(r.Body).Decode(&casePost)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
		} else {
			userID := "1"
			sqlStatement := "INSERT INTO cases (user_id, time_of_symptoms) VALUES ($1, $2) RETURNING case_id;"
			tos := casePost.CaseTime
			var caseID string
			err := db.QueryRow(sqlStatement, userID, tos).Scan(&caseID)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			} else {
				// We now inserted the new covid case
				// We then find all locations visited by the user since the start of symptoms
				rows, err := db.Query("SELECT location_id FROM contact WHERE user_id=$1 AND contact_time >= $2;", userID, tos)
				if err != nil {
					http.Error(rw, err.Error(), http.StatusInternalServerError)
				} else {
					var locations []string
					for rows.Next() {
						var l string
						rows.Scan(&l)
						locations = append(locations, l)
					}
					// We now have all the locations where visitors could have been infected
					// We then find all the users that visited that location after tos
					var usersAtRisk []model.User
					for _, v := range locations {
						// For each location we found, we grab the userIDs that visited it after tos
						rows, err := db.Query("SELECT user_id FROM contact WHERE location_id=$1 AND contact_time >= $2;", v, tos)
						if err != nil {
							http.Error(rw, err.Error(), http.StatusInternalServerError)
						} else {
							for rows.Next() {
								// For each user, grab his info from the users table and add it to usersAtRisk
								var uid string
								rows.Scan(&uid)
								var user model.User
								err := db.QueryRow("SELECT user_id, email FROM users WHERE user_id=$1;", uid).Scan(&user)
								if err != nil {
									http.Error(rw, err.Error(), http.StatusInternalServerError)
								} else {
									usersAtRisk = append(usersAtRisk, user)
								}
							}
						}
					}
					// Now that we have all the users at risk, we notify them and send the list back
					rw.Header().Set("Content-Type", "application/json")
					json.NewEncoder(rw).Encode(usersAtRisk)
				}
			}
		}
	})
}
