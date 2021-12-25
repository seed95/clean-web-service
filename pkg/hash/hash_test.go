package hash

import (
	"github.com/seed95/clean-web-service/pkg/random"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestCheckPassword(t *testing.T) {
	password := random.String(8)
	hashPassword, err := Password(password)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name     string
		password string
		hash     string
		wantErr  error
	}{
		{
			name:     "correct password",
			password: password,
			hash:     hashPassword,
			wantErr:  nil,
		},
		{
			name:     "incorrect password",
			password: random.String(4),
			hash:     hashPassword,
			wantErr:  bcrypt.ErrMismatchedHashAndPassword,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckPassword(tt.password, tt.hash); err != tt.wantErr {
				t.Fatalf("got error: %v, want error: %v", err, tt.wantErr)
			}
		})
	}

}
