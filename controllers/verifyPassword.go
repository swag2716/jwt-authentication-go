package controllers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(providedPassword string, userPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))

	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintln("email or password is incorrect")
		check = false
	}
	return check, msg
}
