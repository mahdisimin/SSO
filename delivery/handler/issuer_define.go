package handler

import (
	"SSO/keys"
	"SSO/services"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v5"
)

func DefineIssuerHandler(issuer services.IssuerStruct) echo.HandlerFunc {
	return func(c *echo.Context) error {
		return c.JSON(http.StatusOK, issuer)
	}
}

func JWKsHandler(jwks keys.JWKs) echo.HandlerFunc {
	return func(c *echo.Context) error {
		return c.JSON(http.StatusOK, jwks)
	}
}

func LoginPageHandler(tmpl *template.Template) echo.HandlerFunc {
	return func(c *echo.Context) error {
		return c.Render(http.StatusOK, "login", tmpl)
	}

}
