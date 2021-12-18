package derrors

import (
	"errors"
	"github.com/seed95/clean-web-service/build/messages"
	"net/http"
)

//go:generate stringer -type kind

type (
	kind uint

	serverError struct {
		kind    kind
		message string
	}
)

const (
	_ kind = iota
	Invalid
	NotFound
	Unauthorized
	Unexpected
	NotAllowed
)

var (
	httpErrors = map[kind]int{
		Invalid:      http.StatusBadRequest,
		NotFound:     http.StatusNotFound,
		Unauthorized: http.StatusUnauthorized,
		Unexpected:   http.StatusInternalServerError,
		NotAllowed:   http.StatusMethodNotAllowed,
	}
)

func New(kind kind, msg string) error {
	return serverError{
		kind:    kind,
		message: msg,
	}
}

func (e serverError) Error() string {
	return e.message
}

//HttpError convert kind of error to Http status error
//if error type is not serverError return StatusInternalServerError(500)
func HttpError(err error) (string, int) {
	var serverErr serverError
	ok := errors.As(err, &serverErr)
	if !ok {
		return messages.GeneralError, http.StatusInternalServerError
	}

	code, ok := httpErrors[serverErr.kind]
	if !ok {
		return serverErr.message, http.StatusBadRequest
	}

	return serverErr.message, code

}

func As(err error) bool {
	return errors.As(err, &serverError{})
}
