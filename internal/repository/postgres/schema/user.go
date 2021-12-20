package schema

import (
	"github.com/seed95/clean-web-service/internal/models"
	"github.com/seed95/clean-web-service/internal/models/types"
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Username    string `gorm:"unique; not null"`
		Password    string
		Firstname   string
		Lastname    string
		Email       string
		PhoneNumber string
		Role        types.Role
		Gender      types.Gender
	}
)

func (user *User) ConvertModel() *models.User {
	return &models.User{
		ID:          user.Model.ID,
		Username:    user.Username,
		Password:    user.Password,
		Firstname:   user.Firstname,
		Lastname:    user.Lastname,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,
		Gender:      user.Gender,
	}

}

func ConvertUser(user *models.User) *User {
	return &User{
		Model:       gorm.Model{ID: user.ID},
		Username:    user.Username,
		Password:    user.Password,
		Firstname:   user.Firstname,
		Lastname:    user.Lastname,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,
		Gender:      user.Gender,
	}
}
