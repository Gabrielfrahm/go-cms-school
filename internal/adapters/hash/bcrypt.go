package hash

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptAdapter struct {
	salt int
}

func NewBcryptAdapter(salt int) *BcryptAdapter {
	return &BcryptAdapter{
		salt: salt,
	}
}

func (b *BcryptAdapter) Hashed(value string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(value), b.salt)
	return string(hashedBytes), err
}

func (b *BcryptAdapter) CompareHashed(value, hashedValue string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedValue), []byte(value))
	return err == nil, err
}
