package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/iamnotrodger/trackster-api/internal/handler"
	"github.com/jmoiron/sqlx"
)

func initDB() (*sqlx.DB, error) {
	return nil, nil
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

	router := mux.NewRouter().StrictSlash(true)

	//Routes
	router.HandleFunc("/", handler.HomePage).Methods("GET")

	//Contact
	router.Handle("/contact", handler.PostContact(db)).Methods("POST")

	log.Fatal(http.ListenAndServe(port, router))
}
