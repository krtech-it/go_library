package middleware_custom

import (
	"github.com/labstack/echo/v4"
	"go_library/internal/api/dto"
	"go_library/internal/errors"
)

func ErrorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}
		if apiErr, ok := err.(*errors.APIError); ok {
			return c.JSON(apiErr.StatusCode, dto.ErrorResponse{Error: apiErr.Message})
		}
		return err
	}
}
