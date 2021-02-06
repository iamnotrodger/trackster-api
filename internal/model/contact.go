package model

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// Contact struct
type Contact struct {
	ContactID   string    `json:"contact_id" db:"contact_id"`
	UserID      string    `json:"user_id" db:"user_id"`
	LocationID  string    `json:"location_id" db:"location_id"`
	ContactTime time.Time `json:"contact_time" db:"contact_time"`
}

// Insert Concact
func (c *Contact) Insert(db *sqlx.DB) error {
	sqlStatement := "INSERT INTO contact (user_id, location_id, contact_time) VALUES ($1, $2, $3) RETURNING contact_id;"

	var contactID string

	err := db.QueryRow(sqlStatement, c.UserID, c.LocationID, c.ContactTime).Scan(&contactID)
	if err != nil {
		return err
	}

	c.ContactID = contactID

	return nil
}
