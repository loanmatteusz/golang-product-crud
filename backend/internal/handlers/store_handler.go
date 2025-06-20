package handlers

import (
	"backend/internal/dtos"
	"backend/internal/services"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type StoreHandler struct {
	service services.StoreService
}

func NewStoreHandler(s services.StoreService) *StoreHandler {
	return &StoreHandler{service: s}
}

func (h *StoreHandler) Create(ctx echo.Context) error {
	var dto dtos.CreateStoreDTO
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	store, err := h.service.Create(dto)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, store)
}

func (h *StoreHandler) GetByID(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid UUID"})
	}

	store, err := h.service.FindByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "store not found"})
	}

	return ctx.JSON(http.StatusOK, store)
}

func (h *StoreHandler) GetAll(ctx echo.Context) error {
	stores, err := h.service.FindAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, stores)
}

func (h *StoreHandler) GetByEstablishmentID(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid UUID"})
	}

	stores, err := h.service.FindByEstablishmentID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, stores)
}

func (h *StoreHandler) Update(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid UUID"})
	}

	var dto dtos.UpdateStoreDTO
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	updated, err := h.service.Update(id, dto)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, updated)
}

func (h *StoreHandler) Delete(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid UUID"})
	}

	if err := h.service.Delete(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.NoContent(http.StatusNoContent)
}
