package handler

import (
	"fmt"
	"net/http"
)

// HomePage handler
func HomePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
	fmt.Fprintf(writer, "Welcome to the Home Page!")
}
