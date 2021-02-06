package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/iamnotrodger/trackster-api/internal/handler"
)

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
	router := mux.NewRouter().StrictSlash(true)

	//Routes
	router.HandleFunc("/", handler.HomePage).Methods("GET")

	log.Fatal(http.ListenAndServe(port, router))
}
