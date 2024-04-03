package routes

import (
	"database/sql"

	"github.com/Gabrielfrahm/go-cms-school/internal/adapters/api/controllers/logout"
	"github.com/Gabrielfrahm/go-cms-school/internal/adapters/repositories"
	usecase "github.com/Gabrielfrahm/go-cms-school/internal/core/usecases/logout"
	"github.com/go-chi/chi/v5"
)

func LogoutRoutes(db *sql.DB) chi.Router {
	r := chi.NewRouter()

	tokenRepo := repositories.NewTokenRepository(db)
	logoutUseCase := usecase.NewLogoutUseCase(*tokenRepo)

	r.Post("/", logout.NewLogoutController(logoutUseCase).Logout) // Login /users
	return r
}
