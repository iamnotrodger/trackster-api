package model

import (
	"github.com/jmoiron/sqlx"
)

// User Struct
type User struct {
	UserID      string `json:"user_id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func (this *User) Insert(db *sqlx.DB) error {
	sqlStatement := "INSERT INTO user (email, phone_number) VALUES ($1, $2) RETURNING user_id;"
	var userID string
	err := db.QueryRow(sqlStatement, this.Email, this.PhoneNumber).Scan(&userID)
	if err != nil {
		return err
	} else {
		this.UserID = userID
		return err
	}
}
