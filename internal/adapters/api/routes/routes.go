package routes

import (
	"github.com/Gabrielfrahm/go-cms-school/internal/adapters/api/controllers/login"
	"github.com/Gabrielfrahm/go-cms-school/internal/adapters/hash"

	usecase "github.com/Gabrielfrahm/go-cms-school/internal/core/usecases/login"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes() *chi.Mux {
	router := chi.NewRouter()

	hasher := hash.NewBcryptAdapter(12)

	loginUseCase := usecase.NewLoginUserCase(nil, hasher)

	router.Mount("/login", LoginRoutes(login.NewLoginController(loginUseCase)))
	return router
}
