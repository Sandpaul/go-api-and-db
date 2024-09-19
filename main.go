package main

import (
	"acme/api"
	"acme/config"
	"acme/db/postgres"
	"acme/repository/user"
	"acme/service"
	"fmt"
	"net/http"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(writer, request)
	})
}

func main() {

	config := config.LoadDataBaseConfig()

	var userRepo user.Repository

	switch config.Type {
	case "postgres":
		connectionString := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", config.User, config.DBName, config.Password, config.Host, config.SSLMode)

		db, err := postgres.PostgresConnection(connectionString)
		if err != nil {
			panic(err)
		}

		userRepo = user.NewPostgresUserRepository(db.DB)

	case "inmemory":
		userRepo = user.NewInMemoryUserRepository()

	default:
		fmt.Errorf("unsupported database type: %s", config.Type)
	}

	userService := service.NewUserService(userRepo)
	userAPI := api.NewUserAPI(userService)

	router := http.NewServeMux()

	router.HandleFunc("GET /", rootHandler)
	router.HandleFunc("GET /api/users", userAPI.GetUsers)
	router.HandleFunc("GET /api/users/{id}", userAPI.GetSingleUser)
	router.HandleFunc("POST /api/users", userAPI.CreateUser)
	router.HandleFunc("DELETE /api/users/{id}", userAPI.DeleteUser)
	router.HandleFunc("PATCH /api/users/{id}", userAPI.UpdateUserName)

	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", CorsMiddleware(router))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, World!")
}
