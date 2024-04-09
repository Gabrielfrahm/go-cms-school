package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	httpError "github.com/Gabrielfrahm/go-cms-school/internal/adapters/api/error"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/usecases"
)

type UserController struct {
	UserUseCase usecases.UserUseCase
}

func NewUserController(userUseCase usecases.UserUseCase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
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

	response, err := u.UserUseCase.CreateUser(userCreateInput)
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

func (u *UserController) ListAllUser(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	var req ListAllUserRequest

	if pageStr := queryParams.Get("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil {
			req.Page = page
		}
	}

	if perPageStr := queryParams.Get("perPage"); perPageStr != "" {
		if perPage, err := strconv.Atoi(perPageStr); err == nil {
			req.PerPage = perPage
		}
	}

	if name := queryParams.Get("name"); name != "" {
		req.Name = &name
	}

	if email := queryParams.Get("email"); email != "" {
		req.Email = &email
	}

	if typeUser := queryParams.Get("type_user"); typeUser != "" {
		req.TypeUser = &typeUser
	}

	if !httpError.ValidateRequest(req, w, CreateUserValidationMessages) {
		return // Pare a execução se a validação falhar
	}

	listUseCaseInput := usecases.ListAllUserInput{
		Page:     &req.Page,
		PerPage:  &req.PerPage,
		Name:     req.Name,
		Email:    req.Email,
		TypeUser: req.TypeUser,
	}

	response, err := u.UserUseCase.ListAllUser(listUseCaseInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ListAllUserResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(ListAllUserResponse))
}
