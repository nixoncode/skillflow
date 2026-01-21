package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Success(c echo.Context, httpCode int, message string, data any) error {
	response := ApiResponse{
		Data:    data,
		Success: true,
		Message: message,
	}
	return c.JSON(httpCode, response)
}

func Error(c echo.Context, httpCode int, message string) error {
	response := ApiResponse{
		Data:    nil,
		Success: false,
		Message: message,
	}
	return c.JSON(httpCode, response)
}

func Ok(c echo.Context, message string, data any) error {
	return Success(c, http.StatusOK, message, data)
}

func Created(c echo.Context, message string, data any) error {
	return Success(c, http.StatusCreated, message, data)
}

func BadRequest(c echo.Context, message string) error {
	return Error(c, http.StatusBadRequest, message)
}

func NotFound(c echo.Context, message string) error {
	return Error(c, http.StatusNotFound, message)
}

func InternalServerError(c echo.Context, message string) error {
	return Error(c, http.StatusInternalServerError, message)
}

func ValidationError(c echo.Context, err error) error {
	return c.JSON(http.StatusUnprocessableEntity, ApiResponse{
		Success: false,
		Message: "validation failed",
		Data:    err,
	})
}
