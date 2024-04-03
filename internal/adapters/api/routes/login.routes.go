package routes

import (
	"database/sql"
	"os"

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
	jwt := jwt.NewJWTAdapter([]byte(os.Getenv("JWT_SECRET")))
	userRepo := repositories.NewUserRepository(db)
	tokenRepo := repositories.NewTokenRepository(db)
	loginUseCase := usecase.NewLoginUserCase(userRepo, tokenRepo, hasher, jwt)

	r.Post("/", login.NewLoginController(loginUseCase).Login) // Login /users
	return r
}
