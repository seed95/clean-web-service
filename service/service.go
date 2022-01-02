package service

import (
	"github.com/seed95/clean-web-service/domain/user"
)

type (
	UserService interface {
		CreateUser(user *user.User) (*user.User, error)
		GetUserById(id uint) (*user.User, error)
		GetUserByUsername(username string) (*user.User, error)
		DeleteUserById(id uint) (*user.User, error)
		DeleteUserByUsername(username string) (*user.User, error)
	}

	ValidationService interface {
		Username(username string) error
		Password(password string) error
	}
)
