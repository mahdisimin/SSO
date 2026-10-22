package handler

import (
	"SSO/services"
	"net/http"

	"github.com/labstack/echo/v5"
)

// DefineIssuerHandler Closure to return a handler func based on configuration
func DefineIssuerHandler(issuer services.IssuerStruct) echo.HandlerFunc {
	return func(c *echo.Context) error {
		return c.JSON(http.StatusOK, issuer)
	}
}
