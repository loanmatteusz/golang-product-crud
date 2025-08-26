package mocks

import (
	"backend/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type CategoryRepository struct {
	mock.Mock
}

func (m *CategoryRepository) Create(category *models.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *CategoryRepository) FindByID(id uuid.UUID) (*models.Category, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Category), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *CategoryRepository) FindByName(name string) (*models.Category, error) {
	args := m.Called(name)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Category), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *CategoryRepository) Update(category *models.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *CategoryRepository) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *CategoryRepository) FindAll(limit, offset int, nameFilter string) ([]models.Category, int64, error) {
	args := m.Called(limit, offset, nameFilter)
	if args.Get(0) != nil {
		return args.Get(0).([]models.Category), args.Get(1).(int64), args.Error(2)
	}
	return nil, 0, args.Error(2)
}
