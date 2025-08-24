package dtos

import "github.com/google/uuid"

type CreateProductDTO struct {
	Name       string     `json:"name" validate:"required,min=1"`
	Price      int        `json:"price" validate:"required,gte=0"`
	CategoryID *uuid.UUID `json:"category_id"`
}

type UpdateProductDTO struct {
	Name       *string    `json:"name" validate:"min=1"`
	Price      *int       `json:"price" validate:"gte=0"`
	CategoryID *uuid.UUID `json:"category_id"`
}

// RESPONSES
type ProductResponseDTO struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	CategoryID uuid.UUID `json:"category_id"`
}
