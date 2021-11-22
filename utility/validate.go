package utility

import "net/mail"

//ValidateEmail user input
func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
