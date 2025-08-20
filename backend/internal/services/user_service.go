package services

import (
	"backend/internal/custom_errors"
	"backend/internal/dtos"
	"backend/internal/repositories"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(dto dtos.CreateUserDTO) (*dtos.UserResponseDTO, error)
	Login(dto dtos.LoginDTO) (*dtos.LoginResponseDTO, error)
	RefreshToken(refreshToken string) (string, error)
}

type userService struct {
	repository repositories.UserRepository
	jwtSecret  string
}

func NewUserService(repository repositories.UserRepository, secret string) UserService {
	return &userService{repository, secret}
}

func (s *userService) Register(dto dtos.CreateUserDTO) (*dtos.UserResponseDTO, error) {
	if existing, _ := s.repository.FindByEmail(dto.Email); existing != nil {
		return nil, custom_errors.ErrEmailAlreadyRegistered
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := dto.ToEntity()
	user.Password = string(hashedPassword)

	if err := s.repository.Create(user); err != nil {
		return nil, err
	}

	return dtos.UserToResponseDTO(user), nil
}

func (s *userService) Login(dto dtos.LoginDTO) (*dtos.LoginResponseDTO, error) {
	user, err := s.repository.FindByEmail(dto.Email)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)) != nil {
		return nil, custom_errors.ErrInvalidCredentials
	}

	accessToken, err := s.generateJWT(user.ID.String(), false)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.generateJWT(user.ID.String(), true)
	if err != nil {
		return nil, err
	}

	return &dtos.LoginResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *userService) RefreshToken(refreshToken string) (string, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, custom_errors.ErrInvalidToken
		}
		return []byte(s.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return "", custom_errors.ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", custom_errors.ErrInvalidToken
	}

	if claims["type"] != "refresh" {
		return "", custom_errors.ErrInvalidToken
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", custom_errors.ErrInvalidToken
	}

	newAccessToken, err := s.generateJWT(userID, false)
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}

// Auxiliar Functions
func (s *userService) generateJWT(userID string, isRefresh bool) (string, error) {
	var exp time.Duration
	tokenType := "access"

	if isRefresh {
		exp = 7 * 24 * time.Hour
		tokenType = "refresh"
	} else {
		exp = 15 * time.Minute
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"type":    tokenType,
		"exp":     time.Now().Add(exp).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}
