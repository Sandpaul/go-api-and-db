package main

import (
	"acme/api"
	"acme/config"
	"acme/db"
	"acme/db/inmemory"
	"acme/db/postgres"
	"acme/service"
	"fmt"
	"io"
	"net/http"
)

func main() {

	config := config.Postgres

	dbRepo, err := initialiseDatabase(config)
	if err != nil {
		fmt.Println("Error initialising the database:", err)
		return
	}
	defer dbRepo.Close()

	userService := service.NewUserService(dbRepo)
	userAPI := api.NewUserAPI(userService)

	router := http.NewServeMux()

	router.HandleFunc("GET /", rootHandler)
	router.HandleFunc("GET /api/users", userAPI.GetUsers)
	router.HandleFunc("GET /api/users/{id}", userAPI.GetSingleUser)
	router.HandleFunc("POST /api/users", userAPI.CreateUser)
	router.HandleFunc("DELETE /api/users/{id}", userAPI.DeleteUser)
	router.HandleFunc("PATCH /api/users/{id}", userAPI.UpdateUserName)

	// Starting the HTTP server on port 8080 and providing router variable to ListenAndServe
	fmt.Println("Server listening on port 8080...")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello, World!")
}

func initialiseDatabase(config config.DatabaseConfig) (db.Repository, error) {
    switch config.Type {
    case "postgres":
        connectionString := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", config.User, config.DBName, config.Password, config.Host, config.SSLMode)
        return postgres.NewPostgresRepository(connectionString)
    case "inmemory":
        return inmemory.NewInMemoryRepository(), nil
    default:
        return nil, fmt.Errorf("unsupported database type: %s", config.Type)
    }
}
