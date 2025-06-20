package handlers

import (
	"backend/internal/dtos"
	"backend/internal/services"

	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type EstablishmentHandler struct {
	service services.EstablishmentService
}

func NewEstablishmentHandler(s services.EstablishmentService) *EstablishmentHandler {
	return &EstablishmentHandler{service: s}
}

func (h *EstablishmentHandler) Create(ctx echo.Context) error {
	var dto dtos.CreateEstablishmentDTO
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	establishment, err := h.service.Create(dto)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, establishment)
}

func (h *EstablishmentHandler) GetByID(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid UUID"})
	}

	establishment, err := h.service.FindByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "establishment not found"})
	}

	return ctx.JSON(http.StatusOK, establishment)
}

func (h *EstablishmentHandler) GetAll(ctx echo.Context) error {
	establishments, err := h.service.FindAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, establishments)
}

func (h *EstablishmentHandler) Update(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid UUID"})
	}

	var dto dtos.UpdateEstablishmentDTO
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	updated, err := h.service.Update(id, dto)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, updated)
}

func (h *EstablishmentHandler) Delete(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid UUID"})
	}

	if err := h.service.Delete(id); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.NoContent(http.StatusNoContent)
}
