package routes

import (
	"database/sql"
	"os"

	"github.com/Gabrielfrahm/go-cms-school/internal/adapters/api/controllers/user"
	"github.com/Gabrielfrahm/go-cms-school/internal/adapters/api/middlewares"
	"github.com/Gabrielfrahm/go-cms-school/internal/adapters/hash"
	"github.com/Gabrielfrahm/go-cms-school/internal/adapters/jwt"
	"github.com/Gabrielfrahm/go-cms-school/internal/adapters/repositories"
	usecase "github.com/Gabrielfrahm/go-cms-school/internal/core/usecases/user"
	"github.com/go-chi/chi/v5"
)

func UserRoutes(db *sql.DB) chi.Router {
	r := chi.NewRouter()
	jwt := jwt.NewJWTAdapter([]byte(os.Getenv("JWT_SECRET")))

	hasher := hash.NewBcryptAdapter(12)
	permissionRepo := repositories.NewProfileRepository(db)
	userRepo := repositories.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(*userRepo, hasher, *permissionRepo)

	r.With(
		middlewares.AuthMiddleware(jwt, db),
		middlewares.PermissionMiddleware(jwt, db, map[string]int{"users": 1}),
	).Post("/", user.NewUserController(userUseCase).CreateUser) // create user /users
	return r
}
