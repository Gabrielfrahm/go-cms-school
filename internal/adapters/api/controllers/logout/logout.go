package logout

import (
	"net/http"
	"strings"

	httpError "github.com/Gabrielfrahm/go-cms-school/internal/adapters/api/error"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/usecases"
)

type LogoutController struct {
	logoutUseCase usecases.LogoutUseCase
}

func NewLogoutController(logoutUseCase usecases.LogoutUseCase) *LogoutController {
	return &LogoutController{
		logoutUseCase: logoutUseCase,
	}
}

func (c *LogoutController) Logout(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header is required", http.StatusUnauthorized)
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	req := &LogoutRequest{
		Token: tokenString,
	}

	if !httpError.ValidateRequest(req, w, validationMessages) {
		return // Pare a execução se a validação falhar
	}

	err := c.logoutUseCase.Logout(req.Token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("log out success"))
}
