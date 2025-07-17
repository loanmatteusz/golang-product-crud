package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name       string    `gorm:"column:name"`
	Price      int       `gorm:"column:price"`
	CategoryID uuid.UUID `gorm:"column:category_id"`
	Category   Category  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Cascade não funciona (F)

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
