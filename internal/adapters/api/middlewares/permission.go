package middlewares

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/adapters"
)

func PermissionMiddleware(jwtAdapter adapters.JWTPort, db *sql.DB, requiredPermissions map[string]int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header is required", http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			userID, err := jwtAdapter.ExtractUserID(tokenString)
			if err != nil {
				http.Error(w, "something is wrong in token", http.StatusUnauthorized)
			}

			rows, err := db.Query(
				`SELECT 
					users, classes, profiles, lessons 
						FROM user_permissions 
							WHERE user_id = $1`, userID,
			)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			// pega as colunas da consulta
			columns, err := rows.Columns()
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			userPermissions := make(map[string]int)
			values := make([]interface{}, len(columns))
			for rows.Next() {
				for i := range values {
					values[i] = new(int)
				}

				err := rows.Scan(values...)
				if err != nil {
					http.Error(w, err.Error(), http.StatusUnauthorized)
					return
				}

				for i, colName := range columns {
					if colName != "user_id" {
						userPermissions[colName] = *(values[i].(*int))
					}
				}
			}

			if err := rows.Err(); err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			defer rows.Close()

			hasPermission := false
			for route, requiredPermission := range requiredPermissions {
				userPermission, exists := userPermissions[route]
				if exists && userPermission >= requiredPermission {
					hasPermission = true
					break
				}
			}

			if hasPermission {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "insufficient permissions", http.StatusForbidden)
			}
			// next.ServeHTTP(w, r)
		})
	}
}
