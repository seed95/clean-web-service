package postgres

import (
	"errors"
	"github.com/seed95/clean-web-service/build/messages"
	"github.com/seed95/clean-web-service/pkg/derrors"
	"gorm.io/gorm"
	"testing"
)

func TestIsErrorNotFound(t *testing.T) {

	tests := []struct {
		name    string
		err     error
		return1 bool
	}{
		{
			name:    "go package error",
			err:     errors.New("record not found"),
			return1: false,
		},
		{
			name:    "derrors",
			err:     derrors.New(derrors.Unexpected, messages.DBError),
			return1: false,
		},
		{
			name:    "gorm error record not found",
			err:     gorm.ErrRecordNotFound,
			return1: true,
		},
		{
			name:    "gorm error invalid db",
			err:     gorm.ErrInvalidDB,
			return1: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if isErrorNotFound(tt.err) != tt.return1 {
				t.Error()
			}
		})
	}

}
