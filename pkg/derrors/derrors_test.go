package derrors

import (
	"errors"
	"net/http"
	"testing"
)

func TestHttpError(t *testing.T) {

	tests := []struct {
		name     string
		err      error
		wantMsg  string
		wantCode int
	}{
		{
			name:     "unauthorized kind derror",
			err:      New(Unauthorized, "Unauthorized"),
			wantMsg:  "Unauthorized",
			wantCode: http.StatusUnauthorized,
		},
		{
			name:     "invalid error",
			err:      errors.New("test http error"),
			wantMsg:  "GeneralError",
			wantCode: http.StatusInternalServerError,
		},
		{
			name:     "nil error",
			err:      nil,
			wantMsg:  "GeneralError",
			wantCode: http.StatusInternalServerError,
		},
		{
			name:     "undefined kind error",
			err:      New(1003, "Unauthorized"),
			wantMsg:  "Unauthorized",
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "empty error message",
			err:      New(Unauthorized, ""),
			wantMsg:  "",
			wantCode: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg, code := HttpError(tt.err)
			if msg != tt.wantMsg {
				t.Errorf("HttpError() gotMessage: %v, wantMessage: %v", msg, tt.wantMsg)
			}
			if code != tt.wantCode {
				t.Errorf("HttpError() gotCode: %v, wantCode: %v", code, tt.wantCode)
			}
		})
	}
}

func TestAs(t *testing.T) {
	var tests = []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "use derror package",
			err:  New(Unexpected, "unexpected"),
			want: true,
		},
		{
			name: "use errors package",
			err:  errors.New("errors package"),
			want: false,
		},
		{
			name: "nil error",
			err:  nil,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := As(tt.err)
			if tt.want != got {
				t.Fatalf("As() got: %v, want: %v", got, tt.want)
			}
		})
	}

}
