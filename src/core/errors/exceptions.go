package errors

import "net/http"


type CustomError struct {
	Status  int 
	Msg     string
	Code    string
	ValidationErrors interface{}
}

func (e *CustomError) Error() string {
	return e.Msg
}

func (e *CustomError) WithValidationErrors(validationErrors interface{}) *CustomError {
	e.ValidationErrors = validationErrors
	return e
}


func New(status int, msg string, code string) *CustomError {
	return &CustomError{
		Status: status,
		Msg:    msg,
		Code:   code,
	}
}


var (
	BadRequestError = func(msg string) *CustomError {
		return New(http.StatusBadRequest, msg, "BAD_REQUEST")
	}
	UnauthorizedError = func(msg string) *CustomError {
		return New(http.StatusUnauthorized, msg, "UNAUTHORIZED")
	}
	InternalServerError = func(msg string) *CustomError {
		return New(http.StatusInternalServerError, msg, "INTERNAL_SERVER_ERROR")
	}
)
