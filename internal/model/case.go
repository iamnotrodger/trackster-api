package model

import "time"

// Case struct
type Case struct {
	CaseID         string    `json:"case_id"`
	UserID         string    `json:"user_id"`
	TimeOfSymptoms time.Time `json:"time_of_symptoms"`
}
