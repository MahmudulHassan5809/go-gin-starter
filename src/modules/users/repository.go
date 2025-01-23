package users

import (
	"errors"
	"gin_starter/src/core/repository"
	"gin_starter/src/core/services"

	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(user *RegisterUserRequest) error
	GetUserByField(field, value string) (*User, error)
	GetAll(filterOptions *services.FilterOptions, entity *User) ([]User, int64, error)
}

type UserRepositoryImpl struct {
	BaseRepo *repository.BaseRepository[User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		BaseRepo: repository.NewBaseRepository[User](db),
	}
}


var validFields = map[string]bool{
	"username": true,
	"email":    true,
}



func (r *UserRepositoryImpl) RegisterUser(user *RegisterUserRequest) error {
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


func (r *UserRepositoryImpl) GetUserByField(field, value string) (*User, error) {
	if !validFields[field] {
		return nil, errors.New("invalid field name")
	}
	return r.BaseRepo.GetByField(field, value)
}

func (r *UserRepositoryImpl) GetAll(filterOptions *services.FilterOptions, entity *User) ([]User, int64, error) {
	return r.BaseRepo.PaginateAndFilter(filterOptions, entity)

}