package auth

import (
	"errors"
	"gin_starter/src/core/common"
	"gin_starter/src/core/security"
	"gin_starter/src/modules/users"
)

type AuthService interface {
	RegisterUser(req *users.CreateUserRequest) error
	LoginUser(req *LoginRequest) (*Tokens, error)
}

type authService struct {
	UserRepo users.UserRepository
	PasswordHasher security.PasswordHandler
	JwtHandler security.JWTHandler
}

func NewAuthService(repo users.UserRepository, PasswordHasher security.PasswordHandler, JwtHandler security.JWTHandler) AuthService {
	return &authService{UserRepo: repo, PasswordHasher: PasswordHasher}
}


func (s *authService) LoginUser(req *LoginRequest) (*Tokens, error) {
	user, err := s.UserRepo.GetUserByField("username", req.Username)
	if err != nil || user == nil {
		return nil, errors.New("invalid username or password")
	}

	if !s.PasswordHasher.VerifyPassword(user.Password, req.Password) {
		return nil, errors.New("invalid username or password")
	}

	accessToken, err := security.JWTHandler.Encode(s.JwtHandler, "access", common.AccessTokenPayload{
		Username: user.Username,
	})
	if err != nil {
		return nil, err
	}

	refreshToken, err := security.JWTHandler.Encode(s.JwtHandler, "refresh", common.RefreshTokenPayload{
		Username: user.Username,
	})
	if err != nil {
		return nil, err
	}

	return &Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) RegisterUser(user *users.CreateUserRequest) error {
	isUserExists, err := s.UserRepo.GetUserByField("username", user.Username)
	if err != nil {
		return err
	}
	if isUserExists != nil {
		return errors.New("user with this username already exists")
	}
	
	hashedPassword, err := s.PasswordHasher.Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.UserRepo.CreateUser(user)
}