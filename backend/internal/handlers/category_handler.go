package handlers

import (
	"backend/internal/dtos"
	"backend/internal/services"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	service services.CategoryService
}

func NewCategoryHandler(s services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: s}
}

func (h *CategoryHandler) Create(ctx echo.Context) error {
	var dto dtos.CreateCategoryDTO
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Error to create a category"})
	}

	category, err := h.service.Create(dto)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, category)
}

func (h *CategoryHandler) GetByID(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	category, err := h.service.FindByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Erro to try to get a category by ID"})
	}

	return ctx.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) GetAll(ctx echo.Context) error {
	categories, err := h.service.FindAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) Update(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var dto dtos.UpdateCategoryDTO
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Error to try to update category"})
	}

	updated, err := h.service.Update(id, dto)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, updated)
}

func (h *CategoryHandler) Delete(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.service.Delete(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Error to try to delete category"})
	}

	return ctx.NoContent(http.StatusNoContent)
}
