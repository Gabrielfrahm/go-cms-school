package repositories

import "github.com/Gabrielfrahm/go-cms-school/internal/core/entities/token"

type ReturnCreate struct {
	Token         string
	Refresh_token string
}

type TokenRepository interface {
	Create(userID string, token *token.Token) (*ReturnCreate, error)
}
