package postgres

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(connectionString string) error {

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("error connecting to the database: %w", err)
	}
	DB = db

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("error pinging the database: %w", err)
	}

	fmt.Println("Successfully connected to the database!")
	return nil
}