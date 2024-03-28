package login

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	user, err := c.loginUseCase.Login(req.Email, req.Password)
	if err != nil {
		fmt.Println("esse aqui é o erro")
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
