package repositories

import (
	"backend/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EstablishmentRepository interface {
	Create(establishment *models.Establishment) error
	FindByID(id uuid.UUID) (*models.Establishment, error)
	Update(establishment *models.Establishment) error
	Delete(id uuid.UUID) error
	FindAll() ([]models.Establishment, error)
	HasLinkedStores(id uuid.UUID) (bool, error)
}

type establishmentRepository struct {
	db *gorm.DB
}

func NewEstablishmentRepository(db *gorm.DB) EstablishmentRepository {
	return &establishmentRepository{db}
}

func (r *establishmentRepository) Create(establishment *models.Establishment) error {
	return r.db.Create(establishment).Error
}

func (r *establishmentRepository) FindByID(id uuid.UUID) (*models.Establishment, error) {
	var establishment models.Establishment
	if err := r.db.Preload("Address").First(&establishment, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &establishment, nil
}

func (r *establishmentRepository) Update(establishment *models.Establishment) error {
	return r.db.Save(establishment).Error
}

func (r *establishmentRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Establishment{}, "id = ?", id).Error
}

func (r *establishmentRepository) FindAll() ([]models.Establishment, error) {
	var establishments []models.Establishment
	err := r.db.Preload("Address").Find(&establishments).Error
	return establishments, err
}

func (r *establishmentRepository) HasLinkedStores(id uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&models.Store{}).Where("establishment_id = ?", id).Count(&count).Error
	return count > 0, err
}
