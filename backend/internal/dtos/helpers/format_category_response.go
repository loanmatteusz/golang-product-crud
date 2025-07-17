package helpers

import (
	"backend/internal/dtos"
	"backend/internal/models"
)

func FromCategoryModel(category *models.Category) dtos.CategoryResponseDTO {
	return dtos.CategoryResponseDTO{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}

func FromCategoryModelList(categories []models.Category) []dtos.CategoryResponseDTO {
	var list []dtos.CategoryResponseDTO
	for _, cat := range categories {
		list = append(list, FromCategoryModel(&cat))
	}
	return list
}
