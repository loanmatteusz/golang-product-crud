package custom_errors

import "errors"

const (
	CodeInvalidParam   = "INVALID_PARAM"
	CodeInvalidInput   = "INVALID_INPUT"
	CodeInvalidToken   = "INVALID_TOKEN"
	CodeInternalServer = "INTERNAL_SERVER_ERROR"
)

var (
	ErrInvalidParam   = errors.New("invalid param or id")
	ErrInvalidInput   = errors.New("invalid input")
	ErrInvalidToken   = errors.New("invalid or expired token")
	ErrInternalServer = errors.New("internal server error")
)
