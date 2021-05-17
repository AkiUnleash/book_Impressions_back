package result

import (
	"github.com/labstack/echo/v4"
)

type APIresponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Json(c echo.Context, status int, message string) error {
	var apierr APIresponse
	apierr.Status = status
	apierr.Message = message
	return c.JSON(status, apierr)
}
