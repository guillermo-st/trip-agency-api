package models

import (
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uint64 `json:"id" db:"user_id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	RoleId    uint64 `json:"role_id" db:"role_id"`
}

func (usr User) MarshalJSON() ([]byte, error) {
	var tmp struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		RoleID    uint64 `json:"role_id"`
	}

	tmp.FirstName = usr.FirstName
	tmp.LastName = usr.LastName
	tmp.Email = usr.Email
	tmp.RoleID = usr.RoleId

	return json.Marshal(&tmp)
}

func (usr *User) HashPassword(password string) error {
	passwdBytes, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	usr.Password = string(passwdBytes)
	return nil
}

func (usr *User) ValidatePassword(attemptedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(attemptedPassword), []byte(usr.Password))
}
