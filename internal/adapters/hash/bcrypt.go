package hash

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptAdapter struct {
	salt int
}

func (b *BcryptAdapter) Hashed(value string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(value), b.salt)
	return string(hashedBytes), err
}

func (b *BcryptAdapter) CompareHashed(value, hashedValue string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedValue), []byte(value))
}
