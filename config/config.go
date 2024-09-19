package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Type     string
	Host     string
	User     string
	Password string
	SSLMode  string
	DBName   string
}

func LoadDataBaseConfig(filename ...string) DatabaseConfig {
	envFile := ".env"
	if len(filename) > 0 {
		envFile = filename[0]
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Println(".env file not found, using environment variables")
	}

	return DatabaseConfig{
		Type:     os.Getenv("DBTYPE"),
		Host:     os.Getenv("DBHOST"),
		User:     os.Getenv("DBUSER"),
		Password: os.Getenv("DBPASSWORD"),
		SSLMode:  os.Getenv("DBSSLMODE"),
		DBName:   os.Getenv("DBNAME"),
	}
}