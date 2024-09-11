package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", rootHandler)

	// Starting the HTTP server on port 8080
	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, "Hello, World!")
}

