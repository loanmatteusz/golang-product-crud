package models

import (
	"time"

	"github.com/google/uuid"
)

type Store struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Number    string    `gorm:"column:number"`
	Name      string    `gorm:"column:name"`
	LegalName string    `gorm:"column:legal_name"`

	AddressID uuid.UUID `gorm:"type:uuid;not null;column:address_id"`
	Address   Address   `gorm:"foreignKey:AddressID"`

	EstablishmentID uuid.UUID `gorm:"type:uuid;not null;column:establishment_id"`

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
