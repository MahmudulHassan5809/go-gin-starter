package users

import "gin_starter/src/core/services"

type UserService interface {
	UserList(filterOptions *services.FilterOptions, entity *User) ([]User, int64, error)
}


type UserServiceImpl struct {
	UserRepo UserRepository
}


func NewUserService(repo UserRepository) UserService {
	return &UserServiceImpl{UserRepo: repo}
}


func (s *UserServiceImpl) UserList(filterOptions *services.FilterOptions, entity *User) ([]User, int64, error) {
	return s.UserRepo.GetAll(filterOptions, entity)
}


