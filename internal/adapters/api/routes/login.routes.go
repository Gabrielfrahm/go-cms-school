package routes

import (
	"database/sql"

	"github.com/Gabrielfrahm/go-cms-school/internal/adapters/api/controllers/login"
	"github.com/Gabrielfrahm/go-cms-school/internal/adapters/hash"
	"github.com/Gabrielfrahm/go-cms-school/internal/adapters/jwt"
	"github.com/Gabrielfrahm/go-cms-school/internal/adapters/repositories"
	usecase "github.com/Gabrielfrahm/go-cms-school/internal/core/usecases/login"
	"github.com/go-chi/chi/v5"
)

func LoginRoutes(db *sql.DB) chi.Router {
	r := chi.NewRouter()

	hasher := hash.NewBcryptAdapter(12)
	jwt := jwt.NewJWTAdapter([]byte("secret"))
	userRepo := repositories.NewUserRepository(db)
	loginUseCase := usecase.NewLoginUserCase(userRepo, hasher, jwt)

	r.Post("/", login.NewLoginController(loginUseCase).Login) // Login /users
	return r
}
