package repositories

import (
	"database/sql"
	"fmt"

	entity "github.com/Gabrielfrahm/go-cms-school/internal/core/entities/token"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/repositories"
)

type TokenRepository struct {
	db *sql.DB
}

func NewTokenRepository(db *sql.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

func (r *TokenRepository) Create(userID string, token *entity.Token) (*repositories.ReturnCreate, error) {
	var response repositories.ReturnCreate
	// check if user has token in db.
	rows, err := r.db.Query(
		"SELECT * FROM user_tokens WHERE user_id = $1", userID,
	)

	if err != nil {
		return &repositories.ReturnCreate{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		fmt.Println("caiu aqui")
		//create token in DB
		statement, err := r.db.Prepare("INSERT INTO user_tokens (user_id, token, refresh_token, expires_at, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING token, refresh_token")
		if err != nil {
			return &repositories.ReturnCreate{}, err
		}
		defer statement.Close()
		err = statement.QueryRow(token.UserID, token.Token, token.RefreshToken, token.ExpiresAt, token.CreatedAt, token.UpdatedAt).Scan(&response.Token, &response.Refresh_token)
		if err != nil {
			return &repositories.ReturnCreate{}, err
		}
	} else {
		fmt.Println("caiu aqui 2")
		statement, err := r.db.Prepare("DELETE FROM user_tokens WHERE user_id = $1")
		if err != nil {
			return &repositories.ReturnCreate{}, err
		}
		defer statement.Close()
		if _, err = statement.Exec(userID); err != nil {
			return &repositories.ReturnCreate{}, err
		}

		statement2, err := r.db.Prepare("INSERT INTO user_tokens (user_id, token, refresh_token, expires_at, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING token, refresh_token")
		if err != nil {
			return &repositories.ReturnCreate{}, err
		}
		defer statement2.Close()
		err = statement2.QueryRow(token.UserID, token.Token, token.RefreshToken, token.ExpiresAt, token.CreatedAt, token.UpdatedAt).Scan(&response.Token, &response.Refresh_token)
		if err != nil {
			return &repositories.ReturnCreate{}, err
		}
	}

	return &response, nil
}
