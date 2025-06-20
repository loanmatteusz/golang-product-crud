package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name     string    `gorm:"column:name"`
	Email    string    `gorm:"column:email"`
	Password string    `gorm:"column:password"`

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
