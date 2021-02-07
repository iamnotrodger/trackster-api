package model

import (
	"github.com/jmoiron/sqlx"
)

// Location struct
type Location struct {
	LocationID string `json:"location_id" db:"location_id"`
	Address    string `json:"address" db:"addr"`
}

// Insert Location
func (l *Location) Insert(db *sqlx.DB) error {
	sqlStatement := "INSERT INTO locations (addr) VALUES ($1) RETURNING location_id;"
	var locationID string
	err := db.QueryRow(sqlStatement, l.Address).Scan(&locationID)

	l.LocationID = locationID

	return err
}

// SelectLocationByAddress location by address
func SelectLocationByAddress(db *sqlx.DB, address string) (Location, error) {
	sqlStatement := "SELECT * FROM locations WHERE addr = $1;"

	var location Location
	err := db.Get(&location, sqlStatement, address)

	return location, err
}
