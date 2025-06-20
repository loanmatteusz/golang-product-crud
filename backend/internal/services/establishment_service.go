package services

import (
	"backend/internal/dtos"
	"backend/internal/models"
	"backend/internal/repositories"
	"fmt"

	"errors"

	"github.com/google/uuid"
)

type EstablishmentService interface {
	Create(dto dtos.CreateEstablishmentDTO) (*models.Establishment, error)
	FindByID(id uuid.UUID) (*models.Establishment, error)
	FindAll() ([]models.Establishment, error)
	Update(id uuid.UUID, dto dtos.UpdateEstablishmentDTO) (*models.Establishment, error)
	Delete(id uuid.UUID) error
}

type establishmentService struct {
	addressRepository       repositories.AddressRepository
	establishmentRepository repositories.EstablishmentRepository
	storeRepository         repositories.StoreRepository
}

func NewEstablishmentService(addressRepository repositories.AddressRepository, establishmentRepository repositories.EstablishmentRepository, storeRepository repositories.StoreRepository) EstablishmentService {
	return &establishmentService{
		addressRepository:       addressRepository,
		establishmentRepository: establishmentRepository,
		storeRepository:         storeRepository,
	}
}

func (s *establishmentService) Create(dto dtos.CreateEstablishmentDTO) (*models.Establishment, error) {
	addressInput := dtos.CreateAddressDTO{
		Street:  dto.Address.Street,
		Number:  dto.Address.Number,
		City:    dto.Address.City,
		State:   dto.Address.State,
		ZipCode: dto.Address.ZipCode,
	}
	address := &models.Address{
		ID:      uuid.New(),
		Street:  addressInput.Street,
		Number:  addressInput.Number,
		City:    addressInput.City,
		State:   addressInput.State,
		ZipCode: addressInput.ZipCode,
	}
	if err := s.addressRepository.Create(address); err != nil {
		return nil, err
	}
	establishment := &models.Establishment{
		ID:        uuid.New(),
		Number:    dto.Number,
		Name:      dto.Name,
		LegalName: dto.LegalName,
		AddressID: address.ID,
	}
	if err := s.establishmentRepository.Create(establishment); err != nil {
		return nil, err
	}
	establishmentWithAddress, err := s.establishmentRepository.FindByID(establishment.ID)
	if err != nil {
		return nil, err
	}
	return establishmentWithAddress, err
}

func (s *establishmentService) FindByID(id uuid.UUID) (*models.Establishment, error) {
	return s.establishmentRepository.FindByID(id)
}

func (s *establishmentService) FindAll() ([]models.Establishment, error) {
	return s.establishmentRepository.FindAll()
}

func (s *establishmentService) Update(id uuid.UUID, dto dtos.UpdateEstablishmentDTO) (*models.Establishment, error) {
	establishment, err := s.establishmentRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if dto.Name != nil {
		establishment.Name = *dto.Name
	}
	if dto.Number != nil {
		establishment.Number = *dto.Number
	}
	if dto.LegalName != nil {
		establishment.LegalName = *dto.LegalName
	}

	if dto.Address != nil {
		addr, err := s.addressRepository.FindByID(establishment.AddressID)
		if err != nil {
			return nil, err
		}

		if dto.Address.Street != nil {
			addr.Street = *dto.Address.Street
		}
		if dto.Address.Number != nil {
			addr.Number = *dto.Address.Number
		}
		if dto.Address.City != nil {
			addr.City = *dto.Address.City
		}
		if dto.Address.State != nil {
			addr.State = *dto.Address.State
		}
		if dto.Address.ZipCode != nil {
			addr.ZipCode = *dto.Address.ZipCode
		}

		if err := s.addressRepository.Update(addr); err != nil {
			return nil, err
		}
	}

	if err := s.establishmentRepository.Update(establishment); err != nil {
		return nil, err
	}

	return s.establishmentRepository.FindByID(establishment.ID)
}

func (s *establishmentService) Delete(id uuid.UUID) error {
	hasStores, err := s.establishmentRepository.HasLinkedStores(id)
	if err != nil {
		return err
	}
	if hasStores {
		return errors.New("cannot delete establishment with linked stores")
	}
	establishment, err := s.establishmentRepository.FindByID(id)
	if err != nil {
		return err
	}
	fmt.Println(establishment)
	s.establishmentRepository.Delete(id)
	return s.addressRepository.Delete(establishment.AddressID)
}
