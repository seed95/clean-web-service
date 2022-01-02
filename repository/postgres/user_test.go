package postgres

import (
	"errors"
	"github.com/seed95/clean-web-service/build/messages"
	"github.com/seed95/clean-web-service/domain/user"
	"github.com/seed95/clean-web-service/pkg/derrors"
	"testing"
)

func TestCreateUser(t *testing.T) {
	setupTest(t)
	defer teardownTest()

	user := newUserTest()

	var tests = []struct {
		name string
		arg1 *user.User
		err  error
	}{
		{
			name: "create a user",
			arg1: user,
			err:  nil,
		},
		{
			name: "duplicate username",
			arg1: user,
			err:  derrors.New(derrors.Unexpected, messages.DBError),
		},
		{
			name: "empty username",
			arg1: &user.User{},
			err:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := repositoryTest.CreateUser(tt.arg1)
			if err != tt.err {
				t.Fatal(err)
			}
		})
	}

}

func TestGetUserById(t *testing.T) {

	setupTest(t)
	defer teardownTest()

	userTest, err := repositoryTest.CreateUser(newUserTest())
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name     string
		userId   uint
		wantUser *user.User
		wantErr  error
	}{
		{
			name:     "user not found",
			userId:   userTest.ID + 1,
			wantUser: nil,
			wantErr:  derrors.New(derrors.NotFound, messages.UserNotFound),
		},
		{
			name:     "user found",
			userId:   userTest.ID,
			wantUser: userTest,
			wantErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := repositoryTest.GetUserById(tt.userId)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got error: %v, want error: %v", err, tt.wantErr)
			}
			if user != nil && tt.wantUser != nil && *user != *tt.wantUser {
				t.Errorf("got user: %v, want user: %v", user, tt.wantUser)
			}
			if (user == nil && tt.wantUser != nil) || (user != nil && tt.wantUser == nil) {
				t.Errorf("got user: %v, want user: %v", user, tt.wantUser)
			}
		})
	}
}

func TestGetUserByUsername(t *testing.T) {
	setupTest(t)
	defer teardownTest()

	userTest, err := repositoryTest.CreateUser(newUserTest())
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name     string
		username string
		wantUser *user.User
		wantErr  error
	}{
		{
			name:     "user not found",
			username: userTest.Username + "test",
			wantUser: nil,
			wantErr:  derrors.New(derrors.NotFound, messages.UserNotFound),
		},
		{
			name:     "user found",
			username: userTest.Username,
			wantUser: userTest,
			wantErr:  nil,
		},
		{
			name:     "empty username",
			username: "",
			wantUser: nil,
			wantErr:  derrors.New(derrors.NotFound, messages.UserNotFound),
		},
		{
			name:     "persian username",
			username: "سجاد",
			wantUser: nil,
			wantErr:  derrors.New(derrors.NotFound, messages.UserNotFound),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := repositoryTest.GetUserByUsername(tt.username)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got error: %v, want error: %v", err, tt.wantErr)
			}
			if user != nil && tt.wantUser != nil && *user != *tt.wantUser {
				t.Errorf("got user: %v, want user: %v", user, tt.wantUser)
			}
			if (user == nil && tt.wantUser != nil) || (user != nil && tt.wantUser == nil) {
				t.Errorf("got user: %v, want user: %v", user, tt.wantUser)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {

	setupTest(t)
	defer teardownTest()

	userCreated := newUserTest()
	userNotCreated := newUserTest()

	userCreated, err := repositoryTest.CreateUser(userCreated)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("user not found", func(t *testing.T) {
		wantErr := derrors.New(derrors.NotFound, messages.UserNotFound)
		user, err := repositoryTest.UpdateUser(userNotCreated)

		if !errors.Is(err, wantErr) {
			t.Errorf("got error: %v, want error: %v", err, wantErr)
		}
		if user != nil {
			t.Errorf("got user: %v, want user: nil", user)
		}
	})

	t.Run("user found and update", func(t *testing.T) {
		userTest := *userCreated
		userTest.Firstname = "update first name"
		userTest.Lastname = "update last name"
		userTest.PhoneNumber = "update phone number"
		userTest.Email = "update email"
		user, err := repositoryTest.UpdateUser(&userTest)
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		if user == nil {
			t.Fatalf("got user: nil")
		}
		if *user != userTest {
			t.Fatalf("got user: %v, want user: %v", user, userTest)
		}
		userUpdated, err := repositoryTest.GetUserById(userCreated.ID)
		if err != nil {
			t.Fatal(err)
		}
		if *userUpdated != userTest {
			t.Fatalf("update user: %v, create user: %v", userUpdated, userCreated)
		}
	})

	t.Run("user found but not change any things to update", func(t *testing.T) {
		userTest := *userCreated
		user, err := repositoryTest.UpdateUser(&userTest)
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		if user == nil {
			t.Fatalf("got user: nil")
		}
		if *user != userTest {
			t.Fatalf("got user: %v, want user: %v", user, userTest)
		}
		userUpdated, err := repositoryTest.GetUserById(userCreated.ID)
		if err != nil {
			t.Fatal(err)
		}
		if *userUpdated != userTest {
			t.Fatalf("update user: %v, create user: %v", userUpdated, userCreated)
		}
	})

	t.Run("user found and update username", func(t *testing.T) {
		wantErr := derrors.New(derrors.NotFound, messages.UserNotFound)
		userTest := *userCreated
		userTest.Username = "update username"
		user, err := repositoryTest.UpdateUser(&userTest)
		if !errors.Is(err, wantErr) {
			t.Fatalf("got error: %v, want error: %v", err, wantErr)
		}
		if user != nil {
			t.Fatalf("got user: %v", user)
		}
	})

}

func TestDeleteUser(t *testing.T) {

	setupTest(t)
	defer teardownTest()

	userCreated := newUserTest()
	userNotCreated := newUserTest()

	userCreated, err := repositoryTest.CreateUser(userCreated)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("user not found", func(t *testing.T) {
		wantErr := derrors.New(derrors.NotFound, messages.UserNotFound)
		err := repositoryTest.DeleteUser(userNotCreated)
		if !errors.Is(err, wantErr) {
			t.Fatalf("got error: %v, want error: %v", err, wantErr)
		}
	})

	t.Run("delete user", func(t *testing.T) {
		if err := repositoryTest.DeleteUser(userCreated); err != nil {
			t.Fatalf("got error: %v, want error: nil", err)
		}

	})

}

func TestIsUsernameExist(t *testing.T) {

	setupTest(t)
	defer teardownTest()

	userNotCreated := newUserTest()
	t.Run("username not found", func(t *testing.T) {
		exist, err := repositoryTest.IsUsernameExist(userNotCreated.Username)
		if err != nil {
			t.Fatal(err)
		}
		if exist {
			t.Fatal("Username found if no such username exists")
		}
	})

	userCreated := newUserTest()
	userCreated, err := repositoryTest.CreateUser(userCreated)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("username found", func(t *testing.T) {
		exist, err := repositoryTest.IsUsernameExist(userCreated.Username)
		if err != nil {
			t.Fatal(err)
		}
		if !exist {
			t.Fatal("Username not found")
		}
	})

	userCreatedWithEmptyUsername := newUserTest()
	userCreatedWithEmptyUsername.Username = ""
	userCreatedWithEmptyUsername, err = repositoryTest.CreateUser(userCreatedWithEmptyUsername)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("empty username found", func(t *testing.T) {
		exist, err := repositoryTest.IsUsernameExist(userCreated.Username)
		if err != nil {
			t.Fatal(err)
		}
		if !exist {
			t.Fatal("Username not found")
		}
	})

}
