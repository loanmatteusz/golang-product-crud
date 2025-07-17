package dtos

import (
	"time"

	"github.com/google/uuid"
)

type CreateCategoryDTO struct {
	Name string `json:"name" validate:"required,min=1"`
}

type UpdateCategoryDTO struct {
	Name *string `json:"name" validate:"required,min=1"`
}

type CategoryResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
