package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ConnectionString string
	Port             int
}

func Load() Config {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	requiredVars := []string{"POSTGRES_USER", "POSTGRES_PASSWORD", "POSTGRES_HOST", "POSTGRES_PORT", "POSTGRES_DB"}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			log.Fatalf("Error: Environment variable %s is not set", v)
		}
	}

	port, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		port = 3333
	}

	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	return Config{
		ConnectionString: connectionString,
		Port:             port,
	}
}
