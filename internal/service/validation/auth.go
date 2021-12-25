package validation

import (
	"github.com/seed95/clean-web-service/build/messages"
	"github.com/seed95/clean-web-service/pkg/derrors"
	"github.com/seed95/clean-web-service/pkg/log"
	"unicode"
)

func (h *handler) Username(username string) error {
	if l := len(username); l < h.cfg.UsernameMinLength || l > h.cfg.UsernameMaxLength {
		h.logger.Error(&log.Field{
			Section:  "service.validation",
			Function: "Username",
			Params:   map[string]interface{}{"username": username},
			Message:  h.translator.Translate(messages.InvalidUsernameLength),
		})
		return derrors.New(derrors.Invalid, messages.InvalidUsernameLength)
	}
	return nil
}

func (h *handler) Password(password string) error {
	var number, upper, special bool
	var letters int

	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			h.logger.Error(&log.Field{
				Section:  "service.validation",
				Function: "Password",
				Params:   map[string]interface{}{"password": password},
				Message:  h.translator.Translate(messages.InvalidPassword),
			})

			return derrors.New(derrors.Invalid, messages.InvalidPassword)
		}
	}

	if letters >= h.cfg.PasswordMinLetters && number && upper && special {
		return nil
	}

	h.logger.Error(&log.Field{
		Section:  "service.validation",
		Function: "Password",
		Params:   map[string]interface{}{"password": password},
		Message:  h.translator.Translate(messages.InvalidPassword),
	})

	return derrors.New(derrors.Invalid, messages.InvalidPassword)
}
