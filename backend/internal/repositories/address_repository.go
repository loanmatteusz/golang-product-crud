package repositories

import (
	"backend/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AddressRepository interface {
	Create(address *models.Address) error
	FindByID(id uuid.UUID) (*models.Address, error)
	Update(address *models.Address) error
	Delete(id uuid.UUID) error
}

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{db}
}

func (r *addressRepository) Create(address *models.Address) error {
	return r.db.Create(address).Error
}

func (r *addressRepository) FindByID(id uuid.UUID) (*models.Address, error) {
	var address models.Address
	if err := r.db.First(&address, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &address, nil
}

func (r *addressRepository) Update(address *models.Address) error {
	return r.db.Save(address).Error
}

func (r *addressRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Address{}, "id = ?", id).Error
}
