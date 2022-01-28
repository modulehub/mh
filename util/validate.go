package util

import (
	"net/mail"

	"github.com/pkg/errors"
)

//ValidateEmailFunc validates users input
func ValidateEmailFunc(input string) error {
	if ok := ValidateEmail(input); !ok {
		return errors.New("invalid email")
	}
	return nil
}

//ValidateEmail user input
func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
