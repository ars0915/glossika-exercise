package cGin

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	if len(password) < 6 || len(password) > 16 {
		return false
	}

	hasLower := false
	hasUpper := false
	hasDigit := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsDigit(char):
			hasDigit = true
		case !unicode.IsLetter(char) && !unicode.IsDigit(char):
			hasSpecial = true
		}
	}

	return hasLower && hasUpper && hasDigit && hasSpecial
}
