package main

import (
	"fmt"
	"regexp"
)

var EmailValidator *regexp.Regexp

func init() {
	compiled, err := regexp.Compile(`.+@.+\..+`)
	if err != nil {
		panic("Failed to compile email regex")
	}

	EmailValidator = compiled
	fmt.Println("Email Validator regex compiled successfully")
}

func isValidEmail(email string) bool {
	return EmailValidator.Match([]byte(email))
}

func main() {
	fmt.Println(isValidEmail("invalid-email"))
	fmt.Println(isValidEmail("john@mail.com"))
}
