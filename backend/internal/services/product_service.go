package services

import (
	"backend/internal/custom_errors"
	"backend/internal/dtos"
	"backend/internal/models"
	"backend/internal/repositories"

	"github.com/google/uuid"
)

type ProductService interface {
	Create(dto dtos.CreateProductDTO) (*models.Product, error)
	FindByID(id uuid.UUID) (*models.Product, error)
	FindAll() ([]models.Product, error)
	Update(id uuid.UUID, dto dtos.UpdateProductDTO) (*models.Product, error)
	Delete(id uuid.UUID) error
}

type productService struct {
	productRepository  repositories.ProductRepository
	categoryRepository repositories.CategoryRepository
}

func NewProductService(productRepository repositories.ProductRepository, categoryRepository repositories.CategoryRepository) ProductService {
	return &productService{
		productRepository:  productRepository,
		categoryRepository: categoryRepository,
	}
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
	product, err := s.productRepository.FindByID(id)
	if err != nil {
		return nil, custom_errors.ErrProductNotFound
	}
	return product, nil
}

func (s *productService) FindAll() ([]models.Product, error) {
	return s.productRepository.FindAll()
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
