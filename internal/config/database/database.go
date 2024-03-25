package database

import (
	"database/sql"
	"time"

	"github.com/Gabrielfrahm/go-cms-school/internal/config"
	_ "github.com/lib/pq"
)

func Connection() (*sql.DB, error) {
	// initial config
	cfg := config.Load()

	db, err := sql.Open("postgres", cfg.ConnectionString)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
