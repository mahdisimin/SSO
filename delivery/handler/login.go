package handler

import (
	"SSO/delivery/dto"
	"bytes"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v5"
)

func LoginPageHandler(tmpl *template.Template) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var buf bytes.Buffer
		if err := tmpl.ExecuteTemplate(&buf, "login.html", dto.LoginPage{}); err != nil {
			return err
		}
		return c.HTML(http.StatusOK, buf.String())
	}

}
