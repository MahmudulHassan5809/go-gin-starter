package common

import "github.com/golang-jwt/jwt/v5"

type AccessTokenPayload struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type RefreshTokenPayload struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}