package model

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// User Struct
type User struct {
	UserID     string    `json:"user_id" db:"user_id"`
	Username   string    `json:"username" db:"username"`
	ProviderID string    `json:"id" db:"provider_id"`
	Email      string    `json:"email" db:"email"`
	Picture    string    `json:"picture" db:"picture"`
	Name       string    `json:"name" db:"name"`
	GivenName  string    `json:"given_name" db:"given_name"`
	FamilyName string    `json:"family_name" db:"family_name"`
	Joined     time.Time `json:"joined" db:"joined"`
}

//Insert user
func (u *User) Insert(db *sqlx.DB) error {
	sqlStatement := "INSERT INTO users (provider_id, email, name, given_name, family_name, picture) VALUES ($1, $2,	$3, $4, $5, $6) RETURNING user_id;"

	var userID string
	err := db.QueryRow(sqlStatement, u.ProviderID, u.Email, u.Name, u.GivenName, u.FamilyName, u.Picture).Scan(&userID)

	u.UserID = userID

	return err
}

//SelectUserByProviderID func
func SelectUserByProviderID(db *sqlx.DB, providerID string) (*User, error) {
	sqlStatement := "SELECT user_id, provider_id, email, name, given_name, family_name, picture, joined FROM users WHERE provider_id  = $1;"

	user := User{}
	err := db.Get(&user, sqlStatement, providerID)

	return &user, err
}
