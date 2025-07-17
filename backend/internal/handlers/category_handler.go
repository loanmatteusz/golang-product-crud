package handlers

import (
	"backend/internal/custom_errors"
	"backend/internal/dtos"
	"backend/internal/services"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	service  services.CategoryService
	validate *validator.Validate
}

func NewCategoryHandler(s services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: s, validate: validator.New()}
}

func (h *CategoryHandler) Create(ctx echo.Context) error {
	var dto dtos.CreateCategoryDTO
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Error to create a category"})
	}

	if err := h.validate.Struct(dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"validation_erros": err.Error()})
	}

	category, err := h.service.Create(dto)
	if err != nil {
		if errors.Is(err, custom_errors.ErrCategoryNameExists) {
			return ctx.JSON(http.StatusConflict, map[string]string{"error": err.Error()})
		}
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
		if errors.Is(err, custom_errors.ErrCategoryNotFound) {
			return ctx.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	return ctx.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) GetAll(ctx echo.Context) error {
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(ctx.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	nameFilter := ctx.QueryParam("name")

	categories, total, totalPages, err := h.service.FindAll(page, limit, nameFilter)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": categories,
		"pagination": map[string]interface{}{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"totalPages": totalPages,
		},
	})
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
