package models

import "github.com/seed95/clean-web-service/internal/models/types"

type (
	User struct {
		ID          uint
		Username    string
		Password    string
		Firstname   string
		Lastname    string
		Email       string
		PhoneNumber string
		Gender      types.Gender
		Role        types.Role
	}
)
