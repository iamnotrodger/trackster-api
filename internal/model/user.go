package model

import (
	"github.com/jmoiron/sqlx"
)

// User Struct
type User struct {
	UserID      string `json:"user_id"`
	Email       string `json:"email"`
}

func (this *User) Insert(db *sqlx.DB) error {
	sqlStatement := "INSERT INTO user (email) VALUES ($1) RETURNING user_id;"
	var userID string
	err := db.QueryRow(sqlStatement, this.Email).Scan(&userID)
	if err != nil {
		return err
	} else {
		this.UserID = userID
		return err
	}
}
