package models

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func assertJson(t testing.TB, expected string, got string) {
	t.Helper()
	if expected != got {
		t.Errorf("Json didn't match. Expected %v got %v", expected, got)
	}
}

func getTestUser(t testing.TB) User {
	t.Helper()
	return User{
		Id:        1,
		FirstName: "Bob",
		LastName:  "Bobbinson",
		Email:     "bob_bobbinson@mail.com",
		Password:  "hunter2",
		RoleId:    1,
	}
}

func TestUser(t *testing.T) {

	t.Run("Marshalling a User struct results in correct json", func(t *testing.T) {
		usr := getTestUser(t)
		expectedJson := "{\"first_name\":\"Bob\",\"last_name\":\"Bobbinson\",\"email\":\"bob_bobbinson@mail.com\",\"role_id\":1}"
		usrJson, _ := usr.MarshalJSON()
		assertJson(t, string(usrJson), expectedJson)
	})

	t.Run("Hashing an user's password results in valid bcrypt hash", func(t *testing.T) {
		usr := getTestUser(t)
		plainPassword := usr.Password
		_ = usr.HashPassword()
		err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(plainPassword))

		if err != nil {
			t.Errorf("User's hashed password didn't match the hashed plaintext password")
		}
	})

	t.Run("Validating an user's matching password should not return an error", func(t *testing.T) {
		usr := getTestUser(t)
		plainPassword := usr.Password
		_ = usr.HashPassword()

		err := usr.ValidatePassword(plainPassword)
		if err != nil {
			t.Errorf("Correct plaintext password was incorrectly validated for User")
		}
	})

	t.Run("Validating an user's incorrect password should return an error", func(t *testing.T) {
		usr := getTestUser(t)
		plainPassword := "thisPasswordIsNotCorrect"
		_ = usr.HashPassword()

		err := usr.ValidatePassword(plainPassword)
		if err == nil {
			t.Errorf("Validating an incorrect password for user should have returned an error")
		}
	})
}
