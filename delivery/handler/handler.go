package handler

import (
	"SSO/delivery/dto"
	"SSO/keys"
	"SSO/services"
	"bytes"
	"fmt"
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
		var buf bytes.Buffer
		if err := tmpl.ExecuteTemplate(&buf, "login.html", dto.LoginPage{}); err != nil {
			return err
		}
		fmt.Println(buf.String())
		return c.HTML(http.StatusOK, buf.String())
	}

}
