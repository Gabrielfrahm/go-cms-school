package usecases

import "github.com/Gabrielfrahm/go-cms-school/internal/core/entities/user"

type LoginUseCase interface {
	Login(email, password string) (*user.User, error)
}
