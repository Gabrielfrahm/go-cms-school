package repositories

import "github.com/Gabrielfrahm/go-cms-school/internal/core/entities/user"

type ListAllUserInput struct {
	Page     *int
	PerPage  *int
	Name     *string
	Email    *string
	TypeUser *string
}

type UserRepository interface {
	FindByEmail(email string) (*user.User, error)
	Create(user *user.User) (*user.User, error)
	ListAllUser(input ListAllUserInput) ([]user.User, int, error)
	FindById(input string) (*user.User, error)
}
