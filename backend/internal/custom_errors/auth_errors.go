package custom_errors

import "errors"

const (
	CodeEmailAlreadyRegistered = "EMAIL_ALREADY_REGISTERED"
	CodeUserNotFound           = "USER_NOT_FOUND"
	CodeInvalidCredentials     = "INVALID_CREDENTIALS"
)

var (
	ErrEmailAlreadyRegistered = errors.New("this email is already registered")
	ErrInvalidCredentials     = errors.New("invalid email or password")
)
