package handler

import (
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
)

type CasePost struct {
	CaseTime time.Time `json:"case_time"`
}

func PostCase(db *sqlx.DB) http.Handler {
	return http.HandlerFunc(func (rw http.ResponseWriter, r *http.Request) {
		// TODO: get user id from header
		
	})
}
