package login

import (
	"github.com/Gabrielfrahm/go-cms-school/internal/core/entities/user"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/adapters"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/repositories"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/usecases"
)

type LoginUseCase struct {
	userRepo repositories.UserRepository
	hasher   adapters.Hash
}

func NewLoginUserCase(userRepo repositories.UserRepository, hasher adapters.Hash) usecases.LoginUseCase {
	return &LoginUseCase{
		userRepo: userRepo,
		hasher:   hasher,
	}
}

// Login implements usecases.LoginUseCase.
func (l *LoginUseCase) Login(email string, password string) (*user.User, error) {
	user, err := l.userRepo.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	_, err = l.hasher.CompareHashed(password, *user.Password)

	if err != nil {
		return nil, err
	}

	return user, nil
}
