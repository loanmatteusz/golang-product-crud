package services

import (
	"backend/internal/dtos"
	"backend/internal/models"
	"backend/internal/repositories"

	"github.com/google/uuid"
)

type StoreService interface {
	Create(dto dtos.CreateStoreDTO) (*models.Store, error)
	FindByID(id uuid.UUID) (*models.Store, error)
	FindAll() ([]models.Store, error)
	FindByEstablishmentID(establishmentID uuid.UUID) ([]models.Store, error)
	Update(id uuid.UUID, dto dtos.UpdateStoreDTO) (*models.Store, error)
	Delete(id uuid.UUID) error
}

type storeService struct {
	addressRepository repositories.AddressRepository
	storeRepository   repositories.StoreRepository
}

func NewStoreService(addressRepository repositories.AddressRepository, storeRepository repositories.StoreRepository) StoreService {
	return &storeService{addressRepository, storeRepository}
}

func (s *storeService) Create(dto dtos.CreateStoreDTO) (*models.Store, error) {
	address := &models.Address{
		ID:      uuid.New(),
		Street:  dto.Address.Street,
		Number:  dto.Address.Number,
		City:    dto.Address.City,
		State:   dto.Address.State,
		ZipCode: dto.Address.ZipCode,
	}
	if err := s.addressRepository.Create(address); err != nil {
		return nil, err
	}
	store := &models.Store{
		ID:              uuid.New(),
		Number:          dto.Number,
		Name:            dto.Name,
		LegalName:       dto.LegalName,
		AddressID:       address.ID,
		EstablishmentID: dto.EstablishmentID,
	}
	if err := s.storeRepository.Create(store); err != nil {
		return nil, err
	}
	return s.storeRepository.FindByID(store.ID)
}

func (s *storeService) FindByID(id uuid.UUID) (*models.Store, error) {
	return s.storeRepository.FindByID(id)
}

func (s *storeService) FindAll() ([]models.Store, error) {
	return s.storeRepository.FindAll()
}

func (s *storeService) FindByEstablishmentID(establishmentID uuid.UUID) ([]models.Store, error) {
	return s.storeRepository.FindByEstablishmentID(establishmentID)
}

func (s *storeService) Update(id uuid.UUID, dto dtos.UpdateStoreDTO) (*models.Store, error) {
	store, err := s.storeRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if dto.Number != nil {
		store.Number = *dto.Number
	}
	if dto.Name != nil {
		store.Name = *dto.Name
	}
	if dto.LegalName != nil {
		store.LegalName = *dto.LegalName
	}

	if dto.Address != nil {
		address, err := s.addressRepository.FindByID(store.AddressID)
		if err != nil {
			return nil, err
		}

		if dto.Address.Street != nil {
			address.Street = *dto.Address.Street
		}
		if dto.Address.Number != nil {
			address.Number = *dto.Address.Number
		}
		if dto.Address.City != nil {
			address.City = *dto.Address.City
		}
		if dto.Address.State != nil {
			address.State = *dto.Address.State
		}
		if dto.Address.ZipCode != nil {
			address.ZipCode = *dto.Address.ZipCode
		}

		if err := s.addressRepository.Update(address); err != nil {
			return nil, err
		}
	}

	err = s.storeRepository.Update(store)
	return store, err
}

func (s *storeService) Delete(id uuid.UUID) error {
	return s.storeRepository.Delete(id)
}
