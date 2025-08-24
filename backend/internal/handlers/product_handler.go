package handlers

import (
	"backend/internal/custom_errors"
	"backend/internal/dtos"
	"backend/internal/dtos/helpers"
	"backend/internal/services"
	"errors"
	"strconv"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	service   services.ProductService
	validator *validator.Validate
}

func NewProductHandler(s services.ProductService) *ProductHandler {
	return &ProductHandler{service: s, validator: validator.New()}
}

func (h *ProductHandler) Create(ctx echo.Context) error {
	var dto dtos.CreateProductDTO
	if err := ctx.Bind(&dto); err != nil {
		return helpers.SendError(ctx, http.StatusBadRequest, custom_errors.CodeInvalidInput, custom_errors.ErrInvalidInput)
	}

	if err := h.validator.Struct(dto); err != nil {
		return helpers.SendError(ctx, http.StatusBadRequest, custom_errors.CodeInvalidInput, custom_errors.ErrInvalidInput)
	}

	product, err := h.service.Create(dto)
	if err != nil {
		if errors.Is(err, custom_errors.ErrCategoryNotFound) {
			return helpers.SendError(ctx, http.StatusNotFound, custom_errors.CodeCategoryNotFound, custom_errors.ErrCategoryNotFound)
		}
		return helpers.SendError(ctx, http.StatusInternalServerError, custom_errors.CodeInternalServer, custom_errors.ErrInternalServer)
	}

	return ctx.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetByID(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return helpers.SendError(ctx, http.StatusBadRequest, custom_errors.CodeInvalidParam, custom_errors.ErrInvalidParam)
	}

	product, err := h.service.FindByID(id)
	if err != nil {
		return helpers.SendError(ctx, http.StatusNotFound, custom_errors.CodeProductNotFound, custom_errors.ErrProductNotFound)
	}

	return ctx.JSON(http.StatusOK, product)
}

func (h *ProductHandler) GetAll(ctx echo.Context) error {
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(ctx.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	nameFilter := ctx.QueryParam("name")

	products, total, totalPages, err := h.service.FindAll(page, limit, nameFilter)
	if err != nil {
		return helpers.SendError(ctx, http.StatusInternalServerError, custom_errors.CodeInternalServer, custom_errors.ErrInternalServer)
	}

	response := helpers.FromProductModelList(products)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": response,
		"pagination": map[string]interface{}{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"totalPages": totalPages,
		},
	})
}

func (h *ProductHandler) Update(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return helpers.SendError(ctx, http.StatusBadRequest, custom_errors.CodeInvalidParam, custom_errors.ErrInvalidParam)
	}

	var dto dtos.UpdateProductDTO
	if err := ctx.Bind(&dto); err != nil {
		return helpers.SendError(ctx, http.StatusBadRequest, custom_errors.CodeInvalidInput, custom_errors.ErrInvalidInput)
	}

	updated, err := h.service.Update(id, dto)
	if err != nil {
		if errors.Is(err, custom_errors.ErrProductNotFound) {
			return helpers.SendError(ctx, http.StatusNotFound, custom_errors.CodeProductNotFound, custom_errors.ErrProductNotFound)
		}
		return helpers.SendError(ctx, http.StatusInternalServerError, custom_errors.CodeInternalServer, custom_errors.ErrInternalServer)
	}

	return ctx.JSON(http.StatusOK, updated)
}

func (h *ProductHandler) Delete(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return helpers.SendError(ctx, http.StatusBadRequest, custom_errors.CodeInvalidParam, custom_errors.ErrInvalidParam)
	}

	if err := h.service.Delete(id); err != nil {
		if errors.Is(err, custom_errors.ErrProductNotFound) {
			return helpers.SendError(ctx, http.StatusNotFound, custom_errors.CodeProductNotFound, custom_errors.ErrProductNotFound)
		}
		return helpers.SendError(ctx, http.StatusInternalServerError, custom_errors.CodeInternalServer, custom_errors.ErrInternalServer)
	}

	return ctx.NoContent(http.StatusNoContent)
}
