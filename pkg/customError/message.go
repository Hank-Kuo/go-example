package customError

import "errors"

var (
	ErrBadRequest          = errors.New("Bad request")
	ErrNotFound            = errors.New("Not Found")
	ErrUnauthorized        = errors.New("Unauthorized")
	ErrForbidden           = errors.New("Forbidden")
	ErrNotRequiredFields   = errors.New("No such required fields")
	ErrBadQueryParams      = errors.New("Invalid query params")
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrRequestTimeoutError = errors.New("Request Timeout")
	ErrContextCancel       = errors.New("Context cacnel")
	ErrDeadlineExceeded    = errors.New("Context exceed")
)
