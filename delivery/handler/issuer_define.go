package handler

import (
	"SSO/services"
	"net/http"

	"github.com/labstack/echo/v5"
)

func GetIssuerHander(c *echo.Context) error {
	isser := services.NewIssuer()
	return c.JSON(http.StatusOK, isser)
}
