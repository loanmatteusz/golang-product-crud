package helpers

import (
	"backend/internal/dtos"
	"backend/internal/models"
)

func FromProductModel(product *models.Product) dtos.ProductResponseDTO {
	return dtos.ProductResponseDTO{
		ID:         product.ID,
		Name:       product.Name,
		Price:      product.Price,
		CategoryID: product.CategoryID,
		CreatedAt:  product.CreatedAt,
		UpdatedAt:  product.UpdatedAt,
	}
}

func FromProductModelList(categories []models.Product) []dtos.ProductResponseDTO {
	var list []dtos.ProductResponseDTO
	for _, cat := range categories {
		list = append(list, FromProductModel(&cat))
	}
	return list
}
