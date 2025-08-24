package custom_errors

import "errors"

const (
	CodeProductNotFound = "PRODUCT_NOT_FOUND"
)

var (
	ErrProductNotFound   = errors.New("product not found")
	ErrProductNameExists = errors.New("product name already exists")
	ErrProductNameEmpty  = errors.New("product name cannot be empty")
)
