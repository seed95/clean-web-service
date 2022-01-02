package repository

import (
	"github.com/seed95/clean-web-service/domain/user"
)

type (
	Repository interface {
		UserRepository
	}

	UserRepository interface {
		CreateUser(user *user.User) (*user.User, error)
		GetUserById(id uint) (*user.User, error)
		GetUserByUsername(username string) (*user.User, error)
		UpdateUser(user *user.User) (*user.User, error)
		DeleteUser(user *user.User) error
		IsUsernameExist(username string) (bool, error)
	}
)
