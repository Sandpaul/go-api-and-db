package main

import (
	"acme/db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {

    router := http.NewServeMux()

    router.HandleFunc("GET /", rootHandler)
    router.HandleFunc("GET /api/users", getUsers)
    router.HandleFunc("POST /api/users", createUser)

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

func getUsers(writer http.ResponseWriter, request *http.Request) {
    fmt.Printf("got /api/users request\n")

	users := db.GetUsers()

	usersJSON, errMarshal := json.Marshal(users)
	if errMarshal != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	_, err := writer.Write(usersJSON)
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func createUser(writer http.ResponseWriter, request *http.Request) {

        var user db.User
        err := json.NewDecoder(request.Body).Decode(&user)
        if err != nil {
            fmt.Println("Error decoding request body:", err)
            http.Error(writer, "Bad Request", http.StatusBadRequest)
            return
        }

        id := db.AddUser(user)
        writer.WriteHeader(http.StatusCreated)
        fmt.Fprintf(writer, "User created successfully: %d", id)
}
