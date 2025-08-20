package handlers

import (
	"backend/internal/custom_errors"
	"backend/internal/dtos"
	"backend/internal/dtos/helpers"
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
		return helpers.SendError(ctx, http.StatusBadRequest, custom_errors.CodeInvalidInput, custom_errors.ErrInvalidInput)
	}

	user, err := h.service.Register(createUser)
	if err != nil {
		return helpers.SendError(ctx, http.StatusBadRequest, custom_errors.CodeEmailAlreadyRegistered, err)
	}

	return ctx.JSON(http.StatusCreated, user)
}

func (h *UserHandler) Login(ctx echo.Context) error {
	var login dtos.LoginDTO
	if err := ctx.Bind(&login); err != nil {
		return helpers.SendError(ctx, http.StatusBadRequest, custom_errors.CodeInvalidInput, custom_errors.ErrInvalidInput)
	}

	token, err := h.service.Login(login)
	if err != nil {
		return helpers.SendError(
			ctx,
			http.StatusUnauthorized,
			custom_errors.CodeInvalidCredentials,
			custom_errors.ErrInvalidCredentials,
		)
	}

	return ctx.JSON(http.StatusOK, token)
}

func (h *UserHandler) Refresh(ctx echo.Context) error {
	var req dtos.RefreshRequestDTO
	if err := ctx.Bind(&req); err != nil {
		return helpers.SendError(ctx, http.StatusBadRequest, custom_errors.CodeInvalidInput, custom_errors.ErrInvalidInput)
	}

	newToken, err := h.service.RefreshToken(req.RefreshToken)
	if err != nil {
		return helpers.SendError(ctx, http.StatusUnauthorized, custom_errors.CodeInvalidToken, custom_errors.ErrInvalidToken)
	}

	return ctx.JSON(http.StatusOK, map[string]string{"access_token": newToken})
}
