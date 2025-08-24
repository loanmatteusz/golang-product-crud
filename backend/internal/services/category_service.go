package services

import (
	config_cache "backend/internal/config/cache"
	"backend/internal/custom_errors"
	"backend/internal/dtos"
	"backend/internal/models"
	"backend/internal/repositories"
	"errors"
	"math"
	"strings"
	"time"

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
	cacheService       *CacheService
}

func NewCategoryService(categoryRepository repositories.CategoryRepository, cacheService *CacheService) CategoryService {
	return &categoryService{categoryRepository, cacheService}
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
	cacheKey := config_cache.CacheKeys.CategoryByID(id.String())

	var category *models.Category
	if err := s.cacheService.GetJSON(cacheKey, &category); err == nil {
		return category, nil
	}

	category, err := s.categoryRepository.FindByID(id)
	if category == nil {
		return nil, custom_errors.ErrCategoryNotFound
	}
	if err != nil {
		return nil, err
	}

	_ = s.cacheService.SetJSON(cacheKey, category, time.Minute)

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

	cacheKey := config_cache.CacheKeys.CategoriesAll(page, limit, nameFilter)
	var categories []models.Category
	var total int64

	if err := s.cacheService.GetJSON(cacheKey, &categories); err == nil {
		_ = s.cacheService.GetJSON(cacheKey+":total", &total)
		totalPages := int(math.Ceil(float64(total) / float64(limit)))
		return categories, total, totalPages, nil
	}

	categories, total, err := s.categoryRepository.FindAll(limit, offset, nameFilter)
	if err != nil {
		return nil, 0, 0, err
	}

	_ = s.cacheService.SetJSON(cacheKey, categories, time.Minute)
	_ = s.cacheService.SetJSON(cacheKey+":total", total, time.Minute)

	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	return categories, total, totalPages, nil
}

func (s *categoryService) Update(id uuid.UUID, dto dtos.UpdateCategoryDTO) (*models.Category, error) {
	if strings.TrimSpace(*dto.Name) == "" {
		return nil, custom_errors.ErrCategoryNameEmpty
	}

	category, err := s.categoryRepository.FindByID(id)
	if category == nil {
		return nil, custom_errors.ErrCategoryNotFound
	}
	if err != nil {
		return nil, err
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
	if category == nil {
		return custom_errors.ErrCategoryNotFound
	}
	if err != nil {
		return err
	}

	return s.categoryRepository.Delete(id)
}
