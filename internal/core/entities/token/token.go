package token

import "time"

type Token struct {
	UserID       string    `json:"user_id"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewToken(
	userId, token, refreshToken string,
	expiresAt, createdAt, updatedAt time.Time,
) *Token {
	return &Token{
		UserID:       userId,
		Token:        token,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}
