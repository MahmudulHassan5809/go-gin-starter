package security

import (
	"errors"

	"gin_starter/src/core/common"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)



type JWTHandler struct{}

var (
	secretKey           = os.Getenv("JWT_SECRET")
	algorithm           = os.Getenv("JWT_ALGORITHM")
	accessExpireMinutes = getEnvAsInt("ACCESS_TOKEN_EXPIRE_MINUTES")
	refreshExpireMinutes = getEnvAsInt("REFRESH_TOKEN_EXPIRE_MINUTES")
)


func getEnvAsInt(key string) int {
	valStr := os.Getenv(key)
	val, err := strconv.Atoi(valStr)
	if err != nil {
		log.Fatalf("Invalid integer value for %s: %v", key, err)
	}
	return val
}


func (JWTHandler) Encode(tokenType string, payload interface{}) (string, error) {
	var expireMinutes int
	if tokenType == "access" {
		expireMinutes = accessExpireMinutes
	} else if tokenType == "refresh" {
		expireMinutes = refreshExpireMinutes
	} else {
		return "", errors.New("invalid token type")
	}
	expiration := time.Now().Add(time.Duration(expireMinutes) * time.Minute)
	claims := jwt.MapClaims{}
	switch p := payload.(type) {
	case common.AccessTokenPayload:
		claims["username"] = p.Username
	case common.RefreshTokenPayload:
		claims["username"] = p.Username
	default:
		return "", errors.New("invalid payload type")
	}
	claims["exp"] = expiration.Unix()
	token := jwt.NewWithClaims(jwt.GetSigningMethod(algorithm), claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}


func (JWTHandler) Decode(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		log.Printf("JWT invalid: %v", err)
		return nil, errors.New("invalid token")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token claims")
}


func (JWTHandler) DecodeExpired(tokenString string) (jwt.MapClaims, error) {
	parser := &jwt.Parser{}
	token, _, err := parser.ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, errors.New("failed to decode expired token")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}
	return nil, errors.New("invalid token claims")
}
