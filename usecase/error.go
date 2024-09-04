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
)
