package repositories

import (
	"backend/internal/models"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	FindByID(id uuid.UUID) (*models.Category, error)
	FindByName(name string) (*models.Category, error)
	Update(category *models.Category) error
	Delete(id uuid.UUID) error
	FindAll(limit int, offset int) ([]models.Category, int64, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) FindByID(id uuid.UUID) (*models.Category, error) {
	var category models.Category
	if err := r.db.First(&category, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) FindByName(name string) (*models.Category, error) {
	var category models.Category
	if err := r.db.First(&category, "name = ?", name).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Category{}, "id = ?", id).Error
}

func (r *categoryRepository) FindAll(limit int, offset int) ([]models.Category, int64, error) {
	var total int64
	if err := r.db.Model(&models.Category{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var categories []models.Category
	if err := r.db.Limit(limit).Offset(offset).Find(&categories).Error; err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}
