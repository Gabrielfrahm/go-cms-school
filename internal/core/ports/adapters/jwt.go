package adapters

import "time"

type JWTPort interface {
	Create(userID string, expiresAt *time.Time) (string, error)
	Validate(tokenString string) (bool, error)
	ExtractUserID(tokenString string) (string, error)
}
