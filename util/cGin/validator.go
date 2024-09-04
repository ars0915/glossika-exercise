package cGin

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var passwordRegex = regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[\W_]).{6,16}$`)

func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	return passwordRegex.MatchString(password)
}
