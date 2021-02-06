package model

import "time"

// Contact struct
type Contact struct {
	ContactID  string    `json:"contact_id"`
	UserID     string    `json:"user_id"`
	LocationID string    `json:"location_id"`
	Timestamp  time.Time `json:"timestamp"`
}
