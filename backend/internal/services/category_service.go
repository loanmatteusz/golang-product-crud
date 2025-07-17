package services

import (
	"backend/internal/custom_errors"
	"backend/internal/dtos"
	"backend/internal/models"
	"backend/internal/repositories"
	"errors"
	"math"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryService interface {
	Create(dto dtos.CreateCategoryDTO) (*models.Category, error)
	FindByID(id uuid.UUID) (*models.Category, error)
	FindAll(page, limit int, nameFilter string) ([]models.Category, int64, int, error)
	Update(id uuid.UUID, dto dtos.UpdateCategoryDTO) (*models.Category, error)
	Delete(id uuid.UUID) error
}

type categoryService struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryService(categoryRepository repositories.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository}
}

func (s *categoryService) Create(dto dtos.CreateCategoryDTO) (*models.Category, error) {
	categoryNameAlreadyExist, err := s.categoryRepository.FindByName(dto.Name)
	if err != nil {
		return nil, err
	}
	if categoryNameAlreadyExist != nil {
		return nil, custom_errors.ErrCategoryNameExists
	}

	category := &models.Category{
		ID:   uuid.New(),
		Name: dto.Name,
	}
	if err := s.categoryRepository.Create(category); err != nil {
		return nil, err
	}

	return s.categoryRepository.FindByID(category.ID)
}

func (s *categoryService) FindByID(id uuid.UUID) (*models.Category, error) {
	category, err := s.categoryRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, custom_errors.ErrCategoryNotFound
	}
	return category, nil
}

func (s *categoryService) FindAll(page, limit int, nameFilter string) ([]models.Category, int64, int, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit
	categories, total, err := s.categoryRepository.FindAll(limit, offset, nameFilter)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	return categories, total, totalPages, nil
}

func (s *categoryService) Update(id uuid.UUID, dto dtos.UpdateCategoryDTO) (*models.Category, error) {
	if strings.TrimSpace(*dto.Name) == "" {
		return nil, custom_errors.ErrCategoryNameEmpty
	}

	category, err := s.categoryRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, custom_errors.ErrCategoryNotFound
	}

	if dto.Name != nil && *dto.Name != category.Name {
		existingCategory, err := s.categoryRepository.FindByName(*dto.Name)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if existingCategory != nil && existingCategory.ID != id {
			return nil, custom_errors.ErrCategoryNameExists
		}
		category.Name = *dto.Name
	}

	err = s.categoryRepository.Update(category)
	return category, err
}

func (s *categoryService) Delete(id uuid.UUID) error {
	category, err := s.categoryRepository.FindByID(id)
	if err != nil {
		return err
	}
	if category == nil {
		return custom_errors.ErrCategoryNotFound
	}

	return s.categoryRepository.Delete(id)
}
