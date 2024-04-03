package login

import (
	"encoding/json"
	"net/http"

	httpError "github.com/Gabrielfrahm/go-cms-school/internal/adapters/api/error"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/usecases"
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

	if !httpError.ValidateRequest(req, w, validationMessages) {
		return // Pare a execução se a validação falhar
	}

	response, err := c.loginUseCase.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	loginResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(loginResponse))
}
