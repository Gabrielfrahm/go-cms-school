package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Gabrielfrahm/go-cms-school/internal/adapters/api/routes"
	"github.com/Gabrielfrahm/go-cms-school/internal/config"
	"github.com/Gabrielfrahm/go-cms-school/internal/config/database"
)

func main() {
	// initial config
	cfg := config.Load()

	db, err := database.Connection()
	if err != nil {
		log.Fatalln("aqui", err)
	}
	defer db.Close()
	// initial routes
	router := routes.SetupRoutes(db)
	router.Mount("/", router)

	fmt.Println("server on!")
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router)
}
