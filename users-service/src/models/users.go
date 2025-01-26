package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Joao-lucas-felix/DevBook-gRPC/users-service/src/security"
	"github.com/badoux/checkmail"
)

// User repsents a user in the application
type User struct {
	ID        uint64
	Name      string
	Nick      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Prepare preapre the data of the user to persist in the database, using the validate function to validates the field, and format to format the string fields and encpryt the password
func (u *User) Prepare(step string) error {
	err := u.validate(step)
	if err != nil {
		return err
	}
	if err := u.formatFields(step); err != nil {
		return err
	}
	return nil
}

func (u *User) validate(step string) error {
	if u.Name == "" {
		return errors.New("name is a required field, must not be blank")
	}
	if u.Nick == "" {
		return errors.New("nick is a required field, must not be blank")
	}
	if u.Email == "" {
		return errors.New("email is a required field, must not be blank")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		errorMsg := fmt.Sprintf("malformed email. the email format are invalid: %s", err.Error())
		return errors.New(errorMsg)
	}
	if u.Password == "" && step == "create" {
		return errors.New("password is a rquired field, must not be blank")
	}
	return nil
}
func (u *User) formatFields(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
	// encrypt the password
	if step == "create" {
		passwordHashed, err := security.Hash(u.Password)
		if err != nil {
			return err
		}
		u.Password = string(passwordHashed)
	}
	return nil
}
