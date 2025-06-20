package dtos

import "github.com/google/uuid"

type CreateEstablishmentDTO struct {
	Name      string           `json:"name" validate:"required"`
	Number    string           `json:"number" validate:"required"`
	LegalName string           `json:"legal_name" validate:"required"`
	Address   CreateAddressDTO `json:"address" validate:"required,dive"`
}

type UpdateEstablishmentDTO struct {
	Name      *string           `json:"name"`
	Number    *string           `json:"number"`
	LegalName *string           `json:"legal_name"`
	Address   *UpdateAddressDTO `json:"address"`
}

// RESPONSES

type EmbeddedAddressDTO struct {
	ID      uuid.UUID `json:"id"`
	Street  string    `json:"street"`
	Number  string    `json:"number"`
	City    string    `json:"city"`
	State   string    `json:"state"`
	ZipCode string    `json:"zip_code"`
}

type EstablishmentResponseDTO struct {
	ID        uuid.UUID          `json:"id"`
	Number    string             `json:"number"`
	Name      string             `json:"name"`
	LegalName string             `json:"legal_name"`
	Address   EmbeddedAddressDTO `json:"address"`
}
