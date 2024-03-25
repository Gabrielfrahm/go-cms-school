package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/Gabrielfrahm/go-cms-school/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// go run internal/migrations/main.go -action=up
// go run internal/migrations/main.go -action=down
func main() {
	// initial configs
	cfg := config.Load()

	var action string
	flag.StringVar(&action, "action", "up", "migration action 'up' or 'down'")

	flag.Parse()

	// Caminho para o diretório de migrações
	migrationsPath := filepath.Join("internal", "migrations", "sql", action)

	// Construir a URL de conexão com o banco de dados (ajustar conforme necessário)
	dbURL := cfg.ConnectionString

	// Criar instância de migration
	m, err := migrate.New(
		"file://"+migrationsPath,
		dbURL,
	)
	if err != nil {
		log.Fatalf("Could not create migration: %v", err)
	}

	// Executar ação de migration baseada na flag de linha de comando
	switch action {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to run up migrations: %v", err)
		}

		fmt.Println("Migrations up applied successfully")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to run down migrations: %v", err)
		}

		fmt.Println("Migrations down applied successfully")
	default:
		log.Fatalf("Invalid action: %v", action)
	}
}
