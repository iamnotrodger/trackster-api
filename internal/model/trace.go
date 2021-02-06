package model

import "github.com/jmoiron/sqlx"

// Trace struct
type Trace struct {
	TraceID   string `json:"trace_id"`
	ContactID string `json:"contact_id"`
	UserID    string `json:"user_id"`
}

//Insert Trace
func (this *Trace) Insert(db *sqlx.DB) error {
	sqlStatement := "INSERT INTO trace (contact_id) VALUES ($1, $2) RETURNING trace_id;"
	var traceID string
	err := db.QueryRow(sqlStatement, this.ContactID).Scan(&traceID)

	this.TraceID = traceID
	return err
}
