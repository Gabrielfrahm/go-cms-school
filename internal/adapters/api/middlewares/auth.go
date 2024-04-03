package middlewares

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/adapters"
)

func AuthMiddleware(jwtAdapter adapters.JWTPort, db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header is required", http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			valid, err := jwtAdapter.Validate(tokenString)
			if err != nil || !valid {
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			rows, err := db.Query(
				"SELECT * FROM user_tokens WHERE token = $1", tokenString,
			)
			if err != nil {
				http.Error(w, "some is wrong", http.StatusUnauthorized)
				return
			}
			defer rows.Close()

			if !rows.Next() {
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
