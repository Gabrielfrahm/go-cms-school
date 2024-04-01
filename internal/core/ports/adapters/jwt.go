package adapters

type JWTPort interface {
	Create(userID string) (string, error)
	Validate(tokenString string) (bool, error)
}
