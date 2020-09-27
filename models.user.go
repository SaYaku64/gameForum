// models.user.go

package main

import (
	"errors"
	"strings"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

func registerNewUser(username, password string) error {
	if strings.TrimSpace(password) == "" {
		return errors.New("The password field can't be empty")
	} else if strings.TrimSpace(username) == "" {
		return errors.New("The username field can't be empty")
	} else if !checkUserExist(username) {
		return errors.New("The username isn't available")
	}

	hPass, err := HashString(password)
	if err != nil {
		return err
	}
	u := user{Username: username, Password: hPass}

	addUserToDB(u)

	return nil
}
