package user

import (
	"errors"

	"github.com/Gabrielfrahm/go-cms-school/internal/adapters/repositories"
	ports "github.com/Gabrielfrahm/go-cms-school/internal/core/ports/repositories"

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
		user, _ := u.userRepo.FindByEmail(*input.Email)

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

	profile, err := u.profileRepo.FindById(input.ProfileId)

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

// ListAllUser implements usecases.UserUseCase.
func (u *UserUserCase) ListAllUser(input usecases.ListAllUserInput) (*usecases.ListAllUserOutput, error) {

	repoInput := ports.ListAllUserInput{
		Page:     input.Page,
		PerPage:  input.PerPage,
		Name:     input.Name,
		Email:    input.Email,
		TypeUser: input.TypeUser,
	}
	users, total, err := u.userRepo.ListAllUser(repoInput)

	if err != nil {
		return nil, errors.New(err.Error())
	}
	output := &usecases.ListAllUserOutput{
		Data: struct {
			Users []user.User
		}{
			Users: users,
		},
		Meta: struct {
			Page     int
			PerPage  int
			Total    int
			LastPage int
		}{
			Page:     *input.Page,
			PerPage:  *input.PerPage,
			Total:    total,
			LastPage: (total + *input.PerPage - 1) / *input.PerPage,
		},
	}

	return output, nil
}
