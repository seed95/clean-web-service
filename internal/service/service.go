package service

import "github.com/seed95/clean-web-service/internal/models"

type (
	UserService interface {
		CreateUser(user *models.User) (*models.User, error)
		GetUserById(id uint) (*models.User, error)
		GetUserByUsername(username string) (*models.User, error)
		DeleteUserById(id uint) (*models.User, error)
		DeleteUserByUsername(username string) (*models.User, error)
	}

	ValidationService interface {
		Username(username string) error
		Password(password string) error
	}
)
