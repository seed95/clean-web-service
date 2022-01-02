package postgres

import (
	"errors"
	"gorm.io/gorm"
)

//isErrorNotFound returns true if the error type is gorm.ErrRecordNotFound.
func isErrorNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
