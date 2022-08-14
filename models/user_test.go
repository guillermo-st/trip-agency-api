package models

import (
	"testing"
)

func assertJson(t testing.TB, expected string, got string) {
	t.Helper()
	if expected != got {
		t.Errorf("Json didn't match. Expected %v got %v", expected, got)
	}
}

func TestUser(t *testing.T) {
	t.Run("Marshalling a User struct results in correct json", func(t *testing.T) {
		usr := User{
			Id:        1,
			FirstName: "Bob",
			LastName:  "Bobbinson",
			Email:     "bob_bobbinson@mail.com",
			Password:  "hunter2",
			RoleId:    1,
		}

		expectedJson := "{\"first_name\":\"Bob\",\"last_name\":\"Bobbinson\",\"email\":\"bob_bobbinson@mail.com\",\"role_id\":1}"
		usrJson, _ := usr.MarshalJSON()
		assertJson(t, string(usrJson), expectedJson)
	})
}
