package users

import (
	"errors"
	"gin_starter/src/core/repository"

	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(user *RegisterUserRequest) error
	GetUserByField(field, value string) (*User, error)
}

type userRepository struct {
	BaseRepo *repository.BaseRepository[User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		BaseRepo: repository.NewBaseRepository[User](db),
	}
}


var validFields = map[string]bool{
	"username": true,
	"email":    true,
}



func (r *userRepository) RegisterUser(user *RegisterUserRequest) error {
	newUser := User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		IsAdmin:  false,
	}
	if err := r.BaseRepo.Create(&newUser); err != nil { 
		return err
	}
	return nil
}


func (r *userRepository) GetUserByField(field, value string) (*User, error) {
	if !validFields[field] {
		return nil, errors.New("invalid field name")
	}
	return r.BaseRepo.GetByField(field, value)
}