package users

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository interface {
	CreateUser(user *CreateUserRequest) error
	GetUserByField(field, value string) (*User, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

var validFields = map[string]bool{
	"username": true,
	"email":    true,
}


func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *CreateUserRequest) error {
	_, err := r.db.Exec(
		context.Background(),
		"INSERT INTO users (username, password, email) VALUES ($1, $2, $3)",
		user.Username, user.Password, user.Email,
	)
	return err
}


func (r *userRepository) GetUserByField(field, value string) (*User, error) {
	if !validFields[field] {
		return nil, errors.New("invalid field name")
	}
	query := "SELECT id, username, password, email FROM users WHERE " + field + " = $1 LIMIT 1"
	row := r.db.QueryRow(context.Background(), query, value)
	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, nil 
		}
		return nil, err 
	}

	return &user, nil
}