package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/iamnotrodger/trackster-api/internal/handler"
)

func handleRequest() {
	var port string

	if portEnv, ok := os.LookupEnv("TRACKSTER_API_PORT"); ok {
		port = ":" + portEnv
	} else {
		port = ":8080"
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler.HomePage).Methods("GET")

	fmt.Println(port)
	fmt.Println(os.Getenv("ACCESS_TOKEN_SECRET"))

	log.Fatal(http.ListenAndServe(port, router))
}

func main() {
	handleRequest()
}
