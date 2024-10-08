package usecase

import (
	"net/http"

	"github.com/ars0915/glossika-exercise/util/cGin"
)

var (
	ErrorTaskNotFound = cGin.CustomError{
		Code:     1001,
		HTTPCode: http.StatusNotFound,
		Message:  "Task not found",
	}

	ErrorEmailRegistered = cGin.CustomError{
		Code:     1002,
		HTTPCode: http.StatusConflict,
		Message:  "Email registered",
	}

	ErrorUserNotFound = cGin.CustomError{
		Code:     1003,
		HTTPCode: http.StatusNotFound,
		Message:  "User not found",
	}

	ErrorUserVerificationFailed = cGin.CustomError{
		Code:     1004,
		HTTPCode: http.StatusUnauthorized,
		Message:  "User verification failed",
	}

	ErrorPasswordVerificationFailed = cGin.CustomError{
		Code:     1005,
		HTTPCode: http.StatusUnauthorized,
		Message:  "Password verification failed",
	}

	ErrorUserUnverified = cGin.CustomError{
		Code:     1006,
		HTTPCode: http.StatusForbidden,
		Message:  "User unverified",
	}
)
