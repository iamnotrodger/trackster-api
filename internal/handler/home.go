package handler

import (
	"fmt"
	"net/http"
)

// Home handler
func Home(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to the Home Page!")
}
