package handlers

import (
	"backend/internal/custom_errors"
	"backend/internal/dtos"
	"backend/internal/dtos/helpers"
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
		return helpers.SendError(ctx, http.StatusBadRequest, custom_errors.CodeInvalidInput, custom_errors.ErrInvalidInput)
	}

	if err := h.validate.Struct(dto); err != nil {
		return helpers.SendError(ctx, http.StatusBadRequest, custom_errors.CodeInvalidInput, custom_errors.ErrInvalidInput)
	}

	category, err := h.service.Create(dto)
	if err != nil {
		if errors.Is(err, custom_errors.ErrCategoryNameExists) {
			return helpers.SendError(ctx, http.StatusBadRequest, custom_errors.CodeCategoryNameExists, custom_errors.ErrCategoryNameExists)
		}
		return helpers.SendError(ctx, http.StatusInternalServerError, custom_errors.CodeInternalServer, custom_errors.ErrInternalServer)
	}

	response := helpers.FromCategoryModel(category)
	return ctx.JSON(http.StatusCreated, response)
}

func (h *CategoryHandler) GetByID(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return helpers.SendError(ctx, http.StatusBadRequest, custom_errors.CodeInvalidParam, custom_errors.ErrInvalidParam)
	}

	category, err := h.service.FindByID(id)
	if err != nil {
		if errors.Is(err, custom_errors.ErrCategoryNotFound) {
			return helpers.SendError(ctx, http.StatusNotFound, custom_errors.CodeCategoryNotFound, custom_errors.ErrCategoryNotFound)
		}
		return helpers.SendError(ctx, http.StatusInternalServerError, custom_errors.CodeInternalServer, custom_errors.ErrInternalServer)
	}

	response := helpers.FromCategoryModel(category)
	return ctx.JSON(http.StatusOK, response)
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
		return helpers.SendError(ctx, http.StatusInternalServerError, custom_errors.CodeInternalServer, custom_errors.ErrInternalServer)
	}

	response := helpers.FromCategoryModelList(categories)
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

func (h *CategoryHandler) Update(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return helpers.SendError(ctx, http.StatusBadRequest, custom_errors.CodeInvalidParam, custom_errors.ErrInvalidParam)
	}

	var dto dtos.UpdateCategoryDTO
	if err := ctx.Bind(&dto); err != nil {
		return helpers.SendError(ctx, http.StatusBadRequest, custom_errors.CodeInvalidInput, custom_errors.ErrInvalidInput)
	}

	if err := h.validate.Struct(dto); err != nil {
		return helpers.SendError(ctx, http.StatusBadRequest, custom_errors.CodeInvalidInput, custom_errors.ErrInvalidInput)
	}

	categoryUpdated, err := h.service.Update(id, dto)
	if err != nil {
		if errors.Is(err, custom_errors.ErrCategoryNotFound) {
			return helpers.SendError(ctx, http.StatusNotFound, custom_errors.CodeCategoryNotFound, custom_errors.ErrCategoryNotFound)
		}
		if errors.Is(err, custom_errors.ErrCategoryNameExists) {
			return helpers.SendError(ctx, http.StatusBadRequest, custom_errors.CodeCategoryNameExists, custom_errors.ErrCategoryNameExists)
		}
		return helpers.SendError(ctx, http.StatusInternalServerError, custom_errors.CodeInternalServer, custom_errors.ErrInternalServer)
	}

	response := helpers.FromCategoryModel(categoryUpdated)
	return ctx.JSON(http.StatusOK, response)
}

func (h *CategoryHandler) Delete(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return helpers.SendError(ctx, http.StatusBadRequest, custom_errors.CodeInvalidParam, custom_errors.ErrInvalidParam)
	}

	if err := h.service.Delete(id); err != nil {
		if errors.Is(err, custom_errors.ErrCategoryNotFound) {
			return helpers.SendError(ctx, http.StatusNotFound, custom_errors.CodeCategoryNotFound, custom_errors.ErrCategoryNotFound)
		}
		return helpers.SendError(ctx, http.StatusInternalServerError, custom_errors.CodeInternalServer, custom_errors.ErrInternalServer)
	}

	return ctx.NoContent(http.StatusNoContent)
}
