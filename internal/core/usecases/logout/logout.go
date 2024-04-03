package logout

import (
	"github.com/Gabrielfrahm/go-cms-school/internal/adapters/repositories"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/usecases"
)

type LogoutUseCase struct {
	tokenRepo repositories.TokenRepository
}

func NewLogoutUseCase(tokenRepo repositories.TokenRepository) usecases.LogoutUseCase {
	return &LogoutUseCase{
		tokenRepo: tokenRepo,
	}
}

func (l *LogoutUseCase) Logout(token string) error {
	err := l.tokenRepo.Destroy(token)

	if err != nil {
		return err
	}

	return nil
}
