package custom_errors

import "errors"

var ErrCategoryNameExists = errors.New("category name already exists")
var ErrCategoryNotFound = errors.New("category not found")
