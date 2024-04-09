package user

import (
	"encoding/json"
	"net/http"

	httpError "github.com/Gabrielfrahm/go-cms-school/internal/adapters/api/error"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/usecases"
)

type UserController struct {
	CreateUserUseCase usecases.UserUseCase
}

func NewUserController(createUserUseCase usecases.UserUseCase) *UserController {
	return &UserController{
		CreateUserUseCase: createUserUseCase,
	}
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !httpError.ValidateRequest(req, w, CreateUserValidationMessages) {
		return // Pare a execução se a validação falhar
	}

	//montando a o input.
	userPermissionInput := usecases.Permissions{
		Users:    &req.Permission.Users,
		Classes:  &req.Permission.Classes,
		Profiles: &req.Permission.Profiles,
		Lessons:  &req.Permission.Lessons,
	}

	userCreateInput := usecases.CreateUserInput{
		Name:       req.Name,
		Email:      req.Email,
		Password:   req.Password,
		TypeUser:   req.TypeUser,
		ProfileId:  req.ProfileId,
		Permission: userPermissionInput,
	}

	response, err := u.CreateUserUseCase.CreateUser(userCreateInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	CreateUserResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(CreateUserResponse))
}
