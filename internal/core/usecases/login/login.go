package login

import (
	"errors"

	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/adapters"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/repositories"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/usecases"
)

type LoginUseCase struct {
	userRepo repositories.UserRepository
	hasher   adapters.Hash
	Jwt      adapters.JWTPort
}

func NewLoginUserCase(userRepo repositories.UserRepository, hasher adapters.Hash, jwt adapters.JWTPort) usecases.LoginUseCase {
	return &LoginUseCase{
		userRepo: userRepo,
		hasher:   hasher,
		Jwt:      jwt,
	}
}

// Login implements usecases.LoginUseCase.
func (l *LoginUseCase) Login(email string, password string) (*usecases.LoginResponse, error) {
	user, err := l.userRepo.FindByEmail(email)

	if err != nil {
		return &usecases.LoginResponse{}, err
	}

	_, err = l.hasher.CompareHashed(password, *user.Password)

	if err != nil {
		return &usecases.LoginResponse{}, errors.New("email or password incorrect")
	}

	token, err := l.Jwt.Create(user.ID)
	if err != nil {
		return &usecases.LoginResponse{}, errors.New(err.Error())
	}

	return &usecases.LoginResponse{
		Token: token,
	}, nil
}
