package user

import (
	"errors"

	"github.com/Gabrielfrahm/go-cms-school/internal/adapters/repositories"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/entities/permission"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/entities/user"
	entity "github.com/Gabrielfrahm/go-cms-school/internal/core/entities/user"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/adapters"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/usecases"
)

type UserUserCase struct {
	userRepo    repositories.UserRepository
	hash        adapters.Hash
	profileRepo repositories.ProfileRepository
}

func NewUserUseCase(userRepo repositories.UserRepository, hash adapters.Hash, profileRepo repositories.ProfileRepository) usecases.UserUseCase {
	return &UserUserCase{
		userRepo:    userRepo,
		hash:        hash,
		profileRepo: profileRepo,
	}
}

// CreateUser implements usecases.UserUseCase.
func (u *UserUserCase) CreateUser(input usecases.CreateUserInput) (*entity.User, error) {
	var password string
	if input.Email != nil {
		user, err := u.userRepo.FindByEmail(*input.Email)

		if err != nil {
			return &entity.User{}, err
		}

		if user != nil {
			return &entity.User{}, errors.New("email already in use")
		}
	}

	if input.Password != nil {
		hashedValue, err := u.hash.Hashed(*input.Password)

		if err != nil {
			return &entity.User{}, errors.New(err.Error())
		}
		password = hashedValue
	}

	profile, err := u.profileRepo.FindByEmail(input.ProfileId)

	if err != nil {
		return &entity.User{}, errors.New(err.Error())
	}

	permissions := permission.NewPermission(*input.Permission.Users, *input.Permission.Classes, *input.Permission.Profiles, *input.Permission.Lessons)

	user := user.NewUser(nil, input.Name, input.Email, &password, input.TypeUser, *profile, *permissions, nil, nil, nil)

	userRepo, err := u.userRepo.Create(user)

	if err != nil {
		return &entity.User{}, errors.New(err.Error())
	}

	return userRepo, nil
}
