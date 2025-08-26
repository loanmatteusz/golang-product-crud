package services_test

import (
	"backend/internal/custom_errors"
	"backend/internal/dtos"
	"backend/internal/models"
	"backend/internal/services"
	mocks "backend/tests/mocks/services"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCategoryService_Create(t *testing.T) {
	mockCatRepo := new(mocks.CategoryRepository)
	mockCache := new(mocks.CacheService)
	service := services.NewCategoryService(mockCatRepo, mockCache)

	t.Run("sucess: category created successfuly", func(t *testing.T) {
		resetMock(mockCatRepo)

		dto := dtos.CreateCategoryDTO{Name: "Nova Categoria"}
		category := &models.Category{ID: uuid.New(), Name: dto.Name}

		mockCatRepo.On("FindByName", dto.Name).Return(nil, nil)
		mockCatRepo.On("Create", mock.AnythingOfType("*models.Category")).Return(nil)
		mockCatRepo.On("FindByID", mock.AnythingOfType("uuid.UUID")).Return(category, nil)

		result, err := service.Create(dto)

		assert.NoError(t, err)
		assert.Equal(t, dto.Name, result.Name)

		mockCatRepo.AssertExpectations(t)
	})

	t.Run("failed: category already exist", func(t *testing.T) {
		resetMock(mockCatRepo)

		dto := dtos.CreateCategoryDTO{Name: "Nova Categoria"}

		existingCategory := &models.Category{ID: uuid.New(), Name: "Nova Categoria"}
		mockCatRepo.On("FindByName", "Nova Categoria").Return(existingCategory, nil)

		result, err := service.Create(dto)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, custom_errors.ErrCategoryNameExists, err)

		mockCatRepo.AssertExpectations(t)
		mockCatRepo.AssertNotCalled(t, "Create", mock.Anything)
	})

	t.Run("failed: error to try to find category by name", func(t *testing.T) {
		resetMock(mockCatRepo)

		dto := dtos.CreateCategoryDTO{Name: "Nova Categoria"}

		mockCatRepo.On("FindByName", "Nova Categoria").Return(nil, errors.New("generic error"))

		result, err := service.Create(dto)

		assert.Error(t, err)
		assert.Nil(t, result)

		mockCatRepo.AssertExpectations(t)
		mockCatRepo.AssertNotCalled(t, "Create", mock.Anything)
	})

	t.Run("failed: error to try to create category", func(t *testing.T) {
		resetMock(mockCatRepo)

		dto := dtos.CreateCategoryDTO{Name: "Nova Categoria"}

		mockCatRepo.On("FindByName", "Nova Categoria").Return(nil, nil)
		mockCatRepo.On("Create", mock.AnythingOfType("*models.Category")).Return(nil, errors.New("generic error"))

		result, err := service.Create(dto)

		assert.Error(t, err)
		assert.Nil(t, result)

		mockCatRepo.AssertExpectations(t)
		mockCatRepo.AssertNotCalled(t, "FindByID", mock.Anything)
	})

	t.Run("failed: error to try to find category by id", func(t *testing.T) {
		resetMock(mockCatRepo)

		dto := dtos.CreateCategoryDTO{Name: "Nova Categoria"}

		mockCatRepo.On("FindByName", "Nova Categoria").Return(nil, nil)
		mockCatRepo.On("Create", mock.AnythingOfType("*models.Category")).Return(nil)
		mockCatRepo.On("FindByID", mock.AnythingOfType("uuid.UUID")).Return(nil, errors.New("generic error"))

		result, err := service.Create(dto)

		assert.Error(t, err)
		assert.Nil(t, result)

		mockCatRepo.AssertExpectations(t)
	})
}

func resetMock(m *mocks.CategoryRepository) {
	m.ExpectedCalls = nil
	m.Calls = nil
}
