package helpers

import (
	"backend/internal/dtos"

	"github.com/labstack/echo/v4"
)

func SendError(ctx echo.Context, httpStatus int, code string, err error) error {
	resp := dtos.ErrorResponse{
		Error:   code,
		Message: err.Error(),
	}
	return ctx.JSON(httpStatus, resp)
}
