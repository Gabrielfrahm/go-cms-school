package login

import (
	"errors"
	"time"

	entity "github.com/Gabrielfrahm/go-cms-school/internal/core/entities/token"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/adapters"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/repositories"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/usecases"
)

type LoginUseCase struct {
	userRepo  repositories.UserRepository
	tokenRepo repositories.TokenRepository
	hasher    adapters.Hash
	Jwt       adapters.JWTPort
}

func NewLoginUserCase(userRepo repositories.UserRepository, tokenRepo repositories.TokenRepository, hasher adapters.Hash, jwt adapters.JWTPort) usecases.LoginUseCase {
	return &LoginUseCase{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
		hasher:    hasher,
		Jwt:       jwt,
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

	token, err := l.Jwt.Create(user.ID, (*time.Time)(nil)) // token
	if err != nil {
		return &usecases.LoginResponse{}, errors.New(err.Error())
	}

	expiryTime := time.Now().Add(24 * time.Hour)
	refresh_token, err := l.Jwt.Create(user.ID, &expiryTime)
	if err != nil {
		return &usecases.LoginResponse{}, errors.New(err.Error())
	}

	tokenRepo, err := l.tokenRepo.Create(user.ID, entity.NewToken(user.ID, token, refresh_token, time.Now().Add(5*time.Minute), time.Now(), time.Now()))

	if err != nil {
		return &usecases.LoginResponse{}, errors.New(err.Error())
	}

	return &usecases.LoginResponse{
		Token:         tokenRepo.Token,
		Refresh_token: tokenRepo.Refresh_token,
	}, nil
}
