package custom_errors

import "errors"

const (
	CodeEmailAlreadyRegistered = "EMAIL_ALREADY_REGISTERED"
	CodeUserNotFound           = "USER_NOT_FOUND"
	CodeInvalidCredentials     = "INVALID_CREDENTIALS"
	CodeInvalidToken           = "INVALID_TOKEN"
)

var (
	ErrEmailAlreadyRegistered = errors.New("this email is already registered")
	ErrInvalidCredentials     = errors.New("invalid email or password")
	ErrInvalidToken           = errors.New("invalid or expired token")
)
