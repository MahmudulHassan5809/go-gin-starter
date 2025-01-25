package common

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenPayload struct {
	UserID string    `json:"user_id"`
	Email  string    `json:"email"`
	Exp    time.Time `json:"exp,omitempty"`
	Sub    string    `json:"sub"`
	jwt.RegisteredClaims
}

type RefreshTokenPayload struct {
	UserID string    `json:"user_id"`
	Exp    time.Time `json:"exp,omitempty"`
	Sub    string    `json:"sub"`
	jwt.RegisteredClaims
}