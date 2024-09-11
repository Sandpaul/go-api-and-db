package main

import (
    "acme/db"
    "encoding/json"
	"fmt"
	"net/http"
	"io"
)

func main() {

	http.HandleFunc("/", rootHandler)
    http.HandleFunc("/api/users", getUsers)

	// Starting the HTTP server on port 8080
	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
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
