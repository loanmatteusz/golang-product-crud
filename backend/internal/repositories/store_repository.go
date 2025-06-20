package repositories

import (
	"backend/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StoreRepository interface {
	Create(store *models.Store) error
	FindByID(id uuid.UUID) (*models.Store, error)
	Update(store *models.Store) error
	Delete(id uuid.UUID) error
	FindAll() ([]models.Store, error)
	FindByEstablishmentID(establishmentID uuid.UUID) ([]models.Store, error)
}

type storeRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) StoreRepository {
	return &storeRepository{db}
}

func (r *storeRepository) Create(store *models.Store) error {
	return r.db.Create(store).Error
}

func (r *storeRepository) FindByID(id uuid.UUID) (*models.Store, error) {
	var store models.Store
	if err := r.db.Preload("Address").First(&store, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &store, nil
}

func (r *storeRepository) Update(store *models.Store) error {
	return r.db.Save(store).Error
}

func (r *storeRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Store{}, "id = ?", id).Error
}

func (r *storeRepository) FindAll() ([]models.Store, error) {
	var stores []models.Store
	err := r.db.Preload("Address").Find(&stores).Error
	return stores, err
}

func (r *storeRepository) FindByEstablishmentID(establishmentID uuid.UUID) ([]models.Store, error) {
	var stores []models.Store
	err := r.db.Preload("Address").Where("establishment_id = ?", establishmentID).Find(&stores).Error
	return stores, err
}
