package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iamnotrodger/trackster-api/internal/handler"
)

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler.HomePage).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequest()
}
