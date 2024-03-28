package main

import (
	"fmt"
	"net/http"

	"github.com/Gabrielfrahm/go-cms-school/internal/adapters/api/routes"
	"github.com/Gabrielfrahm/go-cms-school/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// initial config
	cfg := config.Load()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome!"))
	})

	router := routes.SetupRoutes()
	r.Mount("/", router)

	fmt.Println("server on!")
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r)
}
