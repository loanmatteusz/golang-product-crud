package models

import (
	"time"

	"github.com/google/uuid"
)

type Address struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	Street  string    `gorm:"column:street"`
	Number  string    `gorm:"column:number"`
	City    string    `gorm:"column:city"`
	State   string    `gorm:"column:state"`
	ZipCode string    `gorm:"column:zip_code"`

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
