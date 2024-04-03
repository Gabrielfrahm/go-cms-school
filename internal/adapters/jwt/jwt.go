package jwt

import (
	"fmt"
	"time"

	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/adapters"
	"github.com/golang-jwt/jwt"
)

type JWTClaims struct {
	UserID string
	jwt.StandardClaims
}

type JWTAdapter struct {
	secretKey []byte
}

func NewJWTAdapter(secretKey []byte) adapters.JWTPort {
	return &JWTAdapter{
		secretKey: secretKey,
	}
}

func (jwtAdapter *JWTAdapter) Create(userID string, expiresAt *time.Time) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute) // 5 minutes
	if expiresAt != nil {
		fmt.Println(expiresAt)
		expirationTime = *expiresAt
	}

	claims := &JWTClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtAdapter.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (jwtAdapter *JWTAdapter) Validate(tokenString string) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtAdapter.secretKey, nil
	})

	if err != nil {
		return false, err
	}

	if _, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return true, nil
	}

	return false, nil
}
