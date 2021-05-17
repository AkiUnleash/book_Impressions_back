package err

import (
	"github.com/labstack/echo/v4"
)

type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func JsonErr(c echo.Context, status int, message string) error {
	var apierr APIError
	apierr.Status = status
	apierr.Message = message
	return c.JSON(status, apierr)
}
