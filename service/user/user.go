package user

import (
	"github.com/seed95/clean-web-service/build/messages"
	"github.com/seed95/clean-web-service/domain/user"
	"github.com/seed95/clean-web-service/pkg/derrors"
	"github.com/seed95/clean-web-service/pkg/hash"
	"github.com/seed95/clean-web-service/pkg/log"
)

func (h *handler) CreateUser(user *user.User) (*user.User, error) {

	if err := h.validation.Username(user.Username); err != nil {
		return nil, err
	}

	if err := h.validation.Password(user.Password); err != nil {
		return nil, err
	}

	exist, err := h.userRepo.IsUsernameExist(user.Username)
	if err != nil {
		return nil, err
	}
	if exist {
		h.logger.Error(&log.Field{
			Section:  "service.user",
			Function: "CreateUser",
			Params:   map[string]interface{}{"username": user.Username},
			Message:  h.translator.Translate(messages.DuplicateUsername),
		})
		return nil, derrors.New(derrors.Invalid, messages.DuplicateUsername)
	}

	password, err := hash.Password(user.Password)
	if err != nil {
		h.logger.Error(&log.Field{
			Section:  "service.user",
			Function: "CreateUser",
			Params:   map[string]interface{}{"password": user.Password},
			Message:  err.Error(),
		})
		return nil, derrors.New(derrors.Unexpected, messages.GeneralError)
	}

	user.Password = password
	user, err = h.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (h *handler) GetUserById(id uint) (*user.User, error) {
	return h.userRepo.GetUserById(id)
}

func (h *handler) GetUserByUsername(username string) (*user.User, error) {
	return h.userRepo.GetUserByUsername(username)
}

func (h *handler) DeleteUserById(id uint) (*user.User, error) {

	user, err := h.userRepo.GetUserById(id)
	if err != nil {
		return nil, err
	}

	if err := h.userRepo.DeleteUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (h *handler) DeleteUserByUsername(username string) (*user.User, error) {

	user, err := h.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if err := h.userRepo.DeleteUser(user); err != nil {
		return nil, err
	}

	return user, nil
}
