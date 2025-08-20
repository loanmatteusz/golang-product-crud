package custom_errors

import "errors"

const (
	CodeInvalidInput = "INVALID_INPUT"
)

var (
	ErrInvalidInput = errors.New("invalid input")
)
