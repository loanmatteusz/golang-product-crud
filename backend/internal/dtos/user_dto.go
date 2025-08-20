package dtos

import (
	"backend/internal/models"
	"time"

	"github.com/google/uuid"
)

type CreateUserDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginDTO struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginResponseDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshRequestDTO struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func (dto *CreateUserDTO) ToEntity() *models.User {
	return &models.User{
		ID:       uuid.New(),
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func UserToResponseDTO(user *models.User) *UserResponseDTO {
	return &UserResponseDTO{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
