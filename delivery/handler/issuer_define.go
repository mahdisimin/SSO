package handler

import (
	"SSO/services"
	"net/http"

	"github.com/labstack/echo/v5"
)

func DefineIssuerHandler(issuer services.IssuerStruct) echo.HandlerFunc {
	return func(c *echo.Context) error {
		return c.JSON(http.StatusOK, issuer)
	}
}
