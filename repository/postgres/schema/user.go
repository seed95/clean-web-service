package schema

import (
	"github.com/seed95/clean-web-service/domain/user"
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
		Role        user.Role
		Gender      user.Gender
	}
)

func (u *User) ConvertModel() *user.User {
	return &user.User{
		ID:          u.Model.ID,
		Username:    u.Username,
		Password:    u.Password,
		Firstname:   u.Firstname,
		Lastname:    u.Lastname,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		Role:        u.Role,
		Gender:      u.Gender,
	}

}

func ConvertUser(user *user.User) *User {
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
