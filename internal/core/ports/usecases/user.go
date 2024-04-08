package usecases

import "github.com/Gabrielfrahm/go-cms-school/internal/core/entities/user"

type UserUseCase interface {
	CreateUser(input CreateUserInput) (*user.User, error)
}

type Permissions struct {
	Users    *int
	Classes  *int
	Profiles *int
	Lessons  *int
}

type CreateUserInput struct {
	Name       *string
	Email      *string
	Password   *string
	TypeUser   string
	ProfileId  string
	Permission Permissions
}
