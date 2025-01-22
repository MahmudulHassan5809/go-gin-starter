package users

import (
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(user *RegisterUserRequest) error
	GetUserByField(field, value string) (*User, error)
}

type userRepository struct {
	db *gorm.DB
}

var validFields = map[string]bool{
	"username": true,
	"email":    true,
}


func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}


func (r *userRepository) RegisterUser(user *RegisterUserRequest) error {
	newUser := User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		IsAdmin:  false,
	}
	if err := r.db.Create(&newUser).Error; err != nil {
		return err
	}
	return nil
}


func (r *userRepository) GetUserByField(field, value string) (*User, error) {
	if !validFields[field] {
		return nil, errors.New("invalid field name")
	}
	var user User
	err := r.db.Where(field+" = ?", value).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}