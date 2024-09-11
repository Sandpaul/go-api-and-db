package main

import (
	"acme/db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
    "strconv"
)

func main() {

    router := http.NewServeMux()

    router.HandleFunc("GET /", rootHandler)
    router.HandleFunc("GET /api/users", getUsers)
    router.HandleFunc("GET /api/users/{id}", getSingleUser)
    router.HandleFunc("POST /api/users", createUser)
    router.HandleFunc("DELETE /api/users/{id}", deleteUser)
    router.HandleFunc("PATCH /api/users/{id}", updateUserName)

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

func getSingleUser(writer http.ResponseWriter, request *http.Request) {

    idStr := request.PathValue("id")
    
    id, err := strconv.Atoi(idStr)
    if err != nil {
        fmt.Println("Error parsing ID:", err)
        http.Error(writer, "Bad Request", http.StatusBadRequest)
        return
    }

    user := db.GetUser(id)

    json.NewEncoder(writer).Encode(user)
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

func deleteUser(writer http.ResponseWriter, request *http.Request) {
 
    idStr := request.PathValue("id")

    id, err := strconv.Atoi(idStr)

    if err != nil {
        fmt.Println("Error parsing ID:", err)
        http.Error(writer, "Bad Request", http.StatusBadRequest)
        return
    }
    result := db.DeleteUser(id)
    
    if result {
        writer.WriteHeader(http.StatusOK)
        fmt.Fprintf(writer, "User deleted successfully: %d", id)
    }
}

func updateUserName(writer http.ResponseWriter, request *http.Request) {
    idStr := request.PathValue("id")

    id, err := strconv.Atoi(idStr)
    fmt.Println("ID:", id)
    if err != nil {
        fmt.Println("Error parsing ID:", err)
        http.Error(writer, "Bad Request", http.StatusBadRequest)
        return
    }

    var requestBody struct {
        Name string `json:"name"`
    }
    
    err = json.NewDecoder(request.Body).Decode(&requestBody)
    if err != nil {
        http.Error(writer, "Invalid request body", http.StatusBadRequest)
        return
    }

    if db.UpdateUserName(id, requestBody.Name) {
        writer.WriteHeader(http.StatusOK)
        fmt.Fprintf(writer, "User name updated succesfully")
    }
}