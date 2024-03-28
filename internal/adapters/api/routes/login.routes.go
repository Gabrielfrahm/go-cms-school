package routes

import (
	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/controllers"
	"github.com/go-chi/chi/v5"
)

func LoginRoutes(loginController controllers.LoginController) chi.Router {
	r := chi.NewRouter()
	r.Post("/", loginController.Login) // Login /users

	return r
}
