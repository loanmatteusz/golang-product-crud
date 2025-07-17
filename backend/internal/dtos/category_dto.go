package dtos

type CreateCategoryDTO struct {
	Name string `json:"name" validate:"required,min=1"`
}

type UpdateCategoryDTO struct {
	Name *string `json:"name" validate:"required,min=1"`
}
