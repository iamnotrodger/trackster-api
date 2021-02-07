package handler

import (
	"log"
	"net/http"
)

//LoggingMiddleware func
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method + " " + r.RequestURI + " " + r.Host)
		next.ServeHTTP(w, r)
	})
}
