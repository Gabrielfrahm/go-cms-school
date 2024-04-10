package usecases

import "github.com/Gabrielfrahm/go-cms-school/internal/core/entities/user"

type UserUseCase interface {
	CreateUser(input CreateUserInput) (*user.User, error)
	ListAllUser(input ListAllUserInput) (*ListAllUserOutput, error)
	ListById(input string) (*user.User, error)
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

type ListAllUserInput struct {
	Page     *int
	PerPage  *int
	Name     *string
	Email    *string
	TypeUser *string
}

type ListAllUserOutput struct {
	Data struct {
		Users []user.User
	}
	Meta struct {
		Page     int
		PerPage  int
		Total    int
		LastPage int
	}
}
