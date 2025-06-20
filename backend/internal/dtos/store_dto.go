package dtos

import "github.com/google/uuid"

type CreateStoreDTO struct {
	Number          string           `json:"number" validate:"required"`
	Name            string           `json:"name" validate:"required"`
	LegalName       string           `json:"legal_name" validate:"required"`
	Address         CreateAddressDTO `json:"address" validate:"required"`
	EstablishmentID uuid.UUID        `json:"establishment_id" validate:"required"`
}

type UpdateStoreDTO struct {
	Number    *string           `json:"number"`
	Name      *string           `json:"name"`
	LegalName *string           `json:"legal_name"`
	Address   *UpdateAddressDTO `json:"address"`
}
