package auth

import (
	"context"
	"fmt"
	"gin_starter/src/core/cache"
	"gin_starter/src/core/common"
	"gin_starter/src/core/errors"
	"gin_starter/src/core/security"
	"gin_starter/src/modules/users"
)

type AuthService interface {
	RegisterUser(req *users.RegisterUserRequest) error
	LoginUser(req *LoginRequest) (*Tokens, error)
}

type authService struct {
	UserRepo users.UserRepository
	PasswordHasher security.PasswordHandler
	JwtHandler security.JWTHandler
	CacheManager   *cache.CacheManager
}

func NewAuthService(repo users.UserRepository, PasswordHasher security.PasswordHandler, JwtHandler security.JWTHandler, CacheManager *cache.CacheManager) AuthService {
	return &authService{UserRepo: repo, PasswordHasher: PasswordHasher, JwtHandler: JwtHandler, CacheManager: CacheManager}
}


func (s *authService) LoginUser(req *LoginRequest) (*Tokens, error) {
	ctx := context.Background()
	
	user, err := s.UserRepo.GetUserByField("username", req.Username)
	if err != nil || user == nil {
		return nil, errors.UnauthorizedError("invalid username or password")
	}

	if !s.PasswordHasher.VerifyPassword(user.Password, req.Password) {
		return nil, errors.UnauthorizedError("invalid username or password")
	}

	accessToken, err := security.JWTHandler.Encode(s.JwtHandler, "access", common.AccessTokenPayload{
		Username: user.Username,
	})
	if err != nil {
		return nil, errors.InternalServerError("error generating access token")
	}

	refreshToken, err := security.JWTHandler.Encode(s.JwtHandler, "refresh", common.RefreshTokenPayload{
		Username: user.Username,
	})
	if err != nil {
		return nil, errors.InternalServerError("error generating refresh token")
	}

	cacheKey := cache.CacheTag.Format("USER_DATA_%v", fmt.Sprint(user.ID))
	userData := map[string]string{
		"username": req.Username,
		"email":    user.Email,
		"id":       fmt.Sprint(user.ID),
	}
	s.CacheManager.HMSet(ctx, cacheKey, userData)

	return &Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) RegisterUser(user *users.RegisterUserRequest) error {
	isUserExists, err := s.UserRepo.GetUserByField("username", user.Username)
	if err != nil {
		return err
	}
	if isUserExists != nil {
		return errors.BadRequestError("user with this username already exists")
	}
	
	hashedPassword, err := s.PasswordHasher.Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.UserRepo.RegisterUser(user)
}