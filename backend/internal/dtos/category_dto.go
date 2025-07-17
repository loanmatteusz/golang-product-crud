package dtos

type CreateCategoryDTO struct {
	Name string `json:"name" validate:"required"`
}

type UpdateCategoryDTO struct {
	Name *string `json:"name"`
}
