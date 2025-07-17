package custom_errors

import "errors"

var (
	ErrCategoryNotFound   = errors.New("category not found")
	ErrCategoryNameExists = errors.New("category name already exists")
	ErrCategoryNameEmpty  = errors.New("category name cannot be empty")
)
