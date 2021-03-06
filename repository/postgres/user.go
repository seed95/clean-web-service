package postgres

import (
	"github.com/seed95/clean-web-service/build/messages"
	"github.com/seed95/clean-web-service/domain/user"
	"github.com/seed95/clean-web-service/pkg/derrors"
	"github.com/seed95/clean-web-service/pkg/log"
	"github.com/seed95/clean-web-service/repository/postgres/schema"
)

//CreateUser creates a user on the database on the users table
func (r *repository) CreateUser(user *user.User) (*user.User, error) {

	u := schema.ConvertUser(user)
	if err := r.db.Create(u).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "postgres.user",
			Function: "CreateUser",
			Params:   map[string]interface{}{"user": u},
			Message:  err.Error(),
		})

		return nil, derrors.New(derrors.Unexpected, messages.DBError)
	}

	return u.ConvertModel(), nil

}

func (r *repository) GetUserById(id uint) (*user.User, error) {

	user := &schema.User{}
	if err := r.db.Model(&schema.User{}).Where("id = ?", id).First(user).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "postgres.user",
			Function: "GetUserById",
			Params:   map[string]interface{}{"user_id": id},
			Message:  err.Error(),
		})

		if isErrorNotFound(err) {
			return nil, derrors.New(derrors.NotFound, messages.UserNotFound)
		}

		return nil, derrors.New(derrors.Unexpected, messages.DBError)

	}

	return user.ConvertModel(), nil

}

func (r *repository) GetUserByUsername(username string) (*user.User, error) {

	user := &schema.User{}
	if err := r.db.Model(&schema.User{}).Where("username = ?", username).First(user).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "postgres.user",
			Function: "GetUserByUsername",
			Params:   map[string]interface{}{"username": username},
			Message:  err.Error(),
		})

		if isErrorNotFound(err) {
			return nil, derrors.New(derrors.NotFound, messages.UserNotFound)
		}

		return nil, derrors.New(derrors.Unexpected, messages.DBError)

	}

	return user.ConvertModel(), nil

}

func (r *repository) UpdateUser(user *user.User) (*user.User, error) {

	u := schema.ConvertUser(user)

	exist, err := r.IsUsernameExist(user.Username)
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, derrors.New(derrors.NotFound, messages.UserNotFound)
	}

	if err := r.db.Model(&schema.User{}).Where("id = ?", user.ID).Save(u).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "postgres.user",
			Function: "UpdateUser",
			Params:   map[string]interface{}{"user": user},
			Message:  err.Error(),
		})

		return nil, derrors.New(derrors.Unexpected, messages.DBError)
	}

	return u.ConvertModel(), nil

}

func (r *repository) DeleteUser(user *user.User) error {
	u := schema.ConvertUser(user)

	query := r.db.Model(&schema.User{}).Where("id = ?", user.ID).Delete(u)
	if err := query.Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "DeleteUser",
			Params:   map[string]interface{}{"user": user},
			Message:  err.Error(),
		})

		return derrors.New(derrors.Unexpected, messages.DBError)
	}

	if query.RowsAffected != 1 {
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "DeleteUser",
			Params:   map[string]interface{}{"user": u},
			Message:  r.translator.Translate(messages.UserNotFound),
		})

		return derrors.New(derrors.NotFound, messages.UserNotFound)
	}

	return nil
}

func (r *repository) IsUsernameExist(username string) (bool, error) {

	user := &schema.User{}
	if err := r.db.Model(&schema.User{}).Where("username = ?", username).First(user).Error; err != nil {

		if isErrorNotFound(err) {
			return false, nil
		}

		r.logger.Error(&log.Field{
			Section:  "postgres.user",
			Function: "GetUserByUsername",
			Params:   map[string]interface{}{"username": username},
			Message:  err.Error(),
		})

		return false, derrors.New(derrors.Unexpected, messages.DBError)

	}

	return true, nil
}
