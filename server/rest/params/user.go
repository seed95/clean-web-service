package params

import (
	"github.com/seed95/clean-web-service/build/messages"
	"github.com/seed95/clean-web-service/domain/user"
	"github.com/seed95/clean-web-service/pkg/derrors"
)

//CreateUser params
type (
	CreateUserRequest struct {
		Username    string      `json:"username"`
		Password    string      `json:"password"`
		Firstname   string      `json:"first_name"`
		Lastname    string      `json:"last_name"`
		Email       string      `json:"email"`
		PhoneNumber string      `json:"phone_number"`
		Gender      user.Gender `json:"gender"`
		Role        user.Role   `json:"role"`
	}

	CreateUserResponse struct {
		ID          uint        `json:"id"`
		Username    string      `json:"username"`
		FirstName   string      `json:"first_name"`
		LastName    string      `json:"last_name"`
		Email       string      `json:"email"`
		PhoneNumber string      `json:"phone_number"`
		Gender      user.Gender `json:"gender"`
		Role        user.Role   `json:"role"`
	}
)

type (
	GetUserResponse struct {
		ID          uint        `json:"id"`
		Username    string      `json:"username"`
		FirstName   string      `json:"first_name"`
		LastName    string      `json:"last_name"`
		Email       string      `json:"email"`
		PhoneNumber string      `json:"phone_number"`
		Gender      user.Gender `json:"gender"`
		Role        user.Role   `json:"role"`
	}

	DeleteUserResponse struct {
		ID          uint        `json:"id"`
		Username    string      `json:"username"`
		FirstName   string      `json:"first_name"`
		LastName    string      `json:"last_name"`
		Email       string      `json:"email"`
		PhoneNumber string      `json:"phone_number"`
		Gender      user.Gender `json:"gender"`
		Role        user.Role   `json:"role"`
	}
)

func ConvertCreateUserResponse(user *user.User) *CreateUserResponse {
	return &CreateUserResponse{
		ID:          user.ID,
		Username:    user.Username,
		FirstName:   user.Firstname,
		LastName:    user.Lastname,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Gender:      user.Gender,
		Role:        user.Role,
	}
}

func ConvertGetUserResponse(user *user.User) *GetUserResponse {
	return &GetUserResponse{
		Username:    user.Username,
		FirstName:   user.Firstname,
		LastName:    user.Lastname,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Gender:      user.Gender,
		Role:        user.Role,
	}
}

func ConvertDeleteUserResponse(user *user.User) *DeleteUserResponse {
	return &DeleteUserResponse{
		ID:          user.ID,
		Username:    user.Username,
		FirstName:   user.Firstname,
		LastName:    user.Lastname,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Gender:      user.Gender,
		Role:        user.Role,
	}
}

func ConvertRequest(i interface{}) (*user.User, error) {
	switch v := i.(type) {
	case *CreateUserRequest, CreateUserRequest:
		return convertCreateUserRequest(v.(*CreateUserRequest)), nil
	default:
		return nil, derrors.New(derrors.Invalid, messages.ParseQueryError)
	}
}

func convertCreateUserRequest(req *CreateUserRequest) *user.User {
	return &user.User{
		Username:    req.Username,
		Password:    req.Password,
		Firstname:   req.Firstname,
		Lastname:    req.Lastname,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Gender:      req.Gender,
		Role:        req.Role,
	}

}
