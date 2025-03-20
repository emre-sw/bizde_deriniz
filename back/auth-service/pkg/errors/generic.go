package errors

import "net/http"

func NewError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func BadRequest(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func InternalServerError(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func Conflict(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusConflict,
		Message: message,
	}
}
