package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Gabrielfrahm/go-cms-school/internal/config"
	_ "github.com/lib/pq"
)

func main() {
	// initial configs
	cfg := config.Load()

	db, err := sql.Open("postgres", cfg.ConnectionString)
	if err != nil {
		log.Fatalf("error when connection on db: %v", err)
	}
	defer db.Close()
	path := filepath.Join("internal", "seeds", "seed.sql")

	seedData, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("error on read file seed: %v", err)
	}

	_, err = db.Exec(string(seedData))
	if err != nil {
		log.Fatalf("error when execute seed: %v", err)
	}

	fmt.Println("seeds applied with success!")
}
