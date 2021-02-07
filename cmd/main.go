package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/iamnotrodger/trackster-api/internal/auth"
	"github.com/iamnotrodger/trackster-api/internal/handler"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func initDB() (*sqlx.DB, error) {
	databaseURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		return nil, errors.New("DATABASE_URL missing")
	}
	db, err := sqlx.Connect("postgres", databaseURL)
	return db, err
}

func getPort() string {
	var port string

	if portEnv, ok := os.LookupEnv("TRACKSTER_API_PORT"); ok {
		port = ":" + portEnv
	} else {
		port = ":8080"
	}

	return port
}

func main() {
	port := getPort()

	db, err := initDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)

	//Middleware
	router.Use(handler.LoggingMiddleware)

	//Routes
	router.HandleFunc("/", handler.Home).Methods("GET")

	router.Handle("/api/contact", auth.Middleware(handler.PostContact(db))).Methods("POST")

	router.HandleFunc("/api/login", handler.Login).Methods("GET")
	router.HandleFunc("/api/auth/google", handler.GoogleLogin).Methods("GET")
	router.Handle("/api/auth/google/callback", handler.GoogleCallback(db)).Methods("GET")
	router.HandleFunc("/api/auth/refresh-token", handler.RefreshToken).Methods("GET")

	log.Fatal(http.ListenAndServe(port, router))
}
