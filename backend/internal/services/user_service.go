package services

import (
	"backend/internal/dtos"
	"backend/internal/models"
	"backend/internal/repositories"

	"github.com/google/uuid"
)

type UserService interface {
	Register(dto dtos.CreateUserDTO) (*models.User, error)
	Login(email string, password string) (bool, error)
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{repository}
}

func (s *userService) Register(dto dtos.CreateUserDTO) (*models.User, error) {
	user := models.User{
		ID:       uuid.New(),
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
	if err := s.repository.Create(&user); err != nil {
		return nil, err
	}
	return s.repository.FindByEmail(user.Email)
}

func (s *userService) Login(email string, password string) (bool, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if password != user.Password {
		return false, err
	}
	return true, err
}
