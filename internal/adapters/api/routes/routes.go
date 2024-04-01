package routes

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/Gabrielfrahm/go-cms-school/internal/adapters/api/middlewares"
	"github.com/Gabrielfrahm/go-cms-school/internal/adapters/jwt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(db *sql.DB) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	jwt := jwt.NewJWTAdapter([]byte(os.Getenv("JWT_SECRET")))
	router.Mount("/login", LoginRoutes(db)) //login

	router.With(middlewares.AuthMiddleware(jwt)).Get("/private", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is a private endpoint"))
	})

	return router
}
