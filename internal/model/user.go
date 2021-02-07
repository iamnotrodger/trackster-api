package model

import "github.com/jmoiron/sqlx"

// User Struct
type User struct {
	UserID     string `json:"user_id"`
	ProviderID string `json:"id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
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
func SelectUserByProviderID(db *sqlx.DB, ProviderID string) (*User, error) {
	sqlStatement := "SELECT * FROM users WHERE provider_id = $1;"

	var user User
	err := db.Get(&user, sqlStatement, ProviderID)
	return &user, err
}
