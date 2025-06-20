package models

import (
	"time"

	"github.com/google/uuid"
)

type Establishment struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Number    string    `gorm:"column:number"`
	Name      string    `gorm:"column:name"`
	LegalName string    `gorm:"column:legal_name"`
	AddressID uuid.UUID `gorm:"column:address_id"`
	Address   Address   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Cascade não funciona (F)

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
