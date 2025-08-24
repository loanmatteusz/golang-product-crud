package services

import (
	config_cache "backend/internal/config/cache"
	"backend/internal/custom_errors"
	"backend/internal/dtos"
	"backend/internal/models"
	"backend/internal/repositories"
	"math"
	"time"

	"github.com/google/uuid"
)

type ProductService interface {
	Create(dto dtos.CreateProductDTO) (*models.Product, error)
	FindByID(id uuid.UUID) (*models.Product, error)
	FindAll(page, limit int, nameFilter string) ([]models.Product, int64, int, error)
	Update(id uuid.UUID, dto dtos.UpdateProductDTO) (*models.Product, error)
	Delete(id uuid.UUID) error
}

type productService struct {
	productRepository  repositories.ProductRepository
	categoryRepository repositories.CategoryRepository
	cacheService       *CacheService
}

func NewProductService(
	productRepository repositories.ProductRepository,
	categoryRepository repositories.CategoryRepository,
	cacheService *CacheService,
) ProductService {
	return &productService{productRepository, categoryRepository, cacheService}
}

func (s *productService) Create(dto dtos.CreateProductDTO) (*models.Product, error) {
	if _, err := s.categoryRepository.FindByID(*dto.CategoryID); err != nil {
		return nil, custom_errors.ErrCategoryNotFound
	}

	product := &models.Product{
		ID:         uuid.New(),
		Name:       dto.Name,
		Price:      dto.Price,
		CategoryID: *dto.CategoryID,
	}
	if err := s.productRepository.Create(product); err != nil {
		return nil, err
	}
	productCreated, err := s.productRepository.FindByID(product.ID)
	if err != nil {
		return nil, err
	}
	return productCreated, err
}

func (s *productService) FindByID(id uuid.UUID) (*models.Product, error) {
	cacheKey := config_cache.CacheKeys.ProductByID(id.String())
	var product *models.Product
	if err := s.cacheService.GetJSON(cacheKey, product); err == nil {
		return product, nil
	}

	product, err := s.productRepository.FindByID(id)
	if err != nil {
		return nil, custom_errors.ErrProductNotFound
	}
	_ = s.cacheService.SetJSON(cacheKey, product, time.Minute)

	return product, nil
}

func (s *productService) FindAll(page, limit int, nameFilter string) ([]models.Product, int64, int, error) {
	if page < 0 {
		page = 1
	}
	if limit < 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	cacheKey := config_cache.CacheKeys.ProductsAll(page, limit, nameFilter)

	var products []models.Product
	var total int64

	if err := s.cacheService.GetJSON(cacheKey, &products); err == nil {
		_ = s.cacheService.GetJSON(cacheKey+":total", &total)
		totalPages := int(math.Ceil(float64(total) / float64(limit)))
		return products, total, totalPages, nil
	}

	products, total, err := s.productRepository.FindAll(limit, offset, nameFilter)
	if err != nil {
		return nil, 0, 0, err
	}

	_ = s.cacheService.SetJSON(cacheKey, products, time.Minute)
	_ = s.cacheService.SetJSON(cacheKey+":total", total, time.Minute)

	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	return products, total, totalPages, nil
}

func (s *productService) Update(id uuid.UUID, dto dtos.UpdateProductDTO) (*models.Product, error) {
	product, err := s.productRepository.FindByID(id)
	if err != nil {
		return nil, custom_errors.ErrProductNotFound
	}

	if dto.Name != nil {
		product.Name = *dto.Name
	}

	if dto.Price != nil {
		product.Price = *dto.Price
	}

	if dto.CategoryID != nil {
		product.CategoryID = *dto.CategoryID
	}

	if err := s.productRepository.Update(product); err != nil {
		return nil, err
	}

	return s.productRepository.FindByID(product.ID)
}

func (s *productService) Delete(id uuid.UUID) error {
	product, err := s.productRepository.FindByID(id)
	if err != nil {
		return custom_errors.ErrProductNotFound
	}
	if err := s.productRepository.Delete(product.ID); err != nil {
		return err
	}
	return nil
}
