package main

import (
	"acme/api"
	"fmt"
	"io"
	"net/http"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("GET /", rootHandler)
	router.HandleFunc("GET /api/users", api.GetUsers)
	router.HandleFunc("GET /api/users/{id}", api.GetSingleUser)
	router.HandleFunc("POST /api/users", api.CreateUser)
	router.HandleFunc("DELETE /api/users/{id}", api.DeleteUser)
	router.HandleFunc("PATCH /api/users/{id}", api.UpdateUserName)

	// Starting the HTTP server on port 8080 and providing router variable to ListenAndServe
	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello, World!")
}
