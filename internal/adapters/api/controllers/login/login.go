package login

import (
	"encoding/json"
	"fmt"
	"net/http"

	errorValidate "github.com/Gabrielfrahm/go-cms-school/internal/adapters/api/error"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/usecases"
	"github.com/go-playground/validator"
)

type LoginController struct {
	loginUseCase usecases.LoginUseCase
}

func NewLoginController(loginUseCase usecases.LoginUseCase) *LoginController {
	return &LoginController{
		loginUseCase: loginUseCase,
	}
}

func (c *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			details := make(map[string]string)
			for _, fieldError := range validationErrors {
				details[fieldError.Field()] = fieldError.Tag()
			}
			errorValidate.SendValidationErrorResponse(w, "VALIDATION_ERROR", "Validation error occurred", details)
			return
		}
		http.Error(w, "Validation error", http.StatusBadRequest)
		return
	}

	user, err := c.loginUseCase.Login(req.Email, req.Password)
	if err != nil {
		fmt.Println("esse aqui é o erro", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(userJSON))
}
