package repositories

type TokenRepository interface {
	Create(userID string) error
}
