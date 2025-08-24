package custom_errors

import "errors"

const (
	CodeCategoryNameExists = "CATEGORY_ALREADY_EXISTS"
	CodeCategoryNotFound   = "CATEGORY_NOT_FOUND"
)

var (
	ErrCategoryNotFound   = errors.New("category not found")
	ErrCategoryNameExists = errors.New("category name already exists")
	ErrCategoryNameEmpty  = errors.New("category name cannot be empty")
)
