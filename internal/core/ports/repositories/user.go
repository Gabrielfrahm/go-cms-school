package repositories

import "github.com/Gabrielfrahm/go-cms-school/internal/core/entities/user"

type UserRepository interface {
	FindByEmail(email string) (*user.User, error)
	Create(user *user.User) (*user.User, error)
}
