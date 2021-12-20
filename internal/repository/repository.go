package repository

import "github.com/seed95/clean-web-service/internal/models"

type (
	Repository interface {
		UserRepository
	}

	UserRepository interface {
		CreateUser(user *models.User) (*models.User, error)
		GetUserById(id uint) (*models.User, error)
		GetUserByUsername(username string) (*models.User, error)
		UpdateUser(user *models.User) (*models.User, error)
		DeleteUser(user *models.User) error
		IsUsernameExist(username string) (bool, error)
	}
)
