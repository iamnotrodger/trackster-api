package model

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// Case struct
type Case struct {
	CaseID         string    `json:"case_id"`
	UserID         string    `json:"user_id"`
	TimeOfSymptoms time.Time `json:"time_of_symptoms"`
}

//Insert Case
func (c *Case) Insert(db *sqlx.DB) error {
	sqlStatement := "INSERT INTO cases (user_id, time_of_symptoms) VALUES ($1, $2) RETURNING case_id;"

	var caseID string
	err := db.QueryRow(sqlStatement, c.UserID, c.TimeOfSymptoms).Scan(&caseID)

	c.CaseID = caseID
	return err
}
