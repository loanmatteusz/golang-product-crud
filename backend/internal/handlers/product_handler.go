package handlers

import (
	"backend/internal/dtos"
	"backend/internal/services"

	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	service services.ProductService
}

func NewProductHandler(s services.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) Create(ctx echo.Context) error {
	var dto dtos.CreateProductDTO
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Error to create a product"})
	}

	product, err := h.service.Create(dto)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Error to create a product"})
	}

	return ctx.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetByID(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid UUID"})
	}

	product, err := h.service.FindByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "product not found"})
	}

	return ctx.JSON(http.StatusOK, product)
}

func (h *ProductHandler) GetAll(ctx echo.Context) error {
	products, err := h.service.FindAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "error to try to get product list"})
	}

	return ctx.JSON(http.StatusOK, products)
}

func (h *ProductHandler) Update(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid UUID"})
	}

	var dto dtos.UpdateProductDTO
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	updated, err := h.service.Update(id, dto)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, updated)
}

func (h *ProductHandler) Delete(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid UUID"})
	}

	if err := h.service.Delete(id); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.NoContent(http.StatusNoContent)
}
