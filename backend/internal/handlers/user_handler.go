package handlers

import (
	"backend/internal/dtos"
	"backend/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(s services.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) Register(ctx echo.Context) error {
	var createUser dtos.CreateUserDTO
	if err := ctx.Bind(&createUser); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	user, err := h.service.Register(createUser)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, user)
}

func (h *UserHandler) Login(ctx echo.Context) error {
	var login dtos.LoginDTO
	if err := ctx.Bind(&login); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}
	logged, err := h.service.Login(login.Email, login.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, false)
	}
	if !logged {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid credentials"})
	}
	return ctx.JSON(http.StatusOK, logged)
}
