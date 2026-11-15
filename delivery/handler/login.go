package handler

import (
	"SSO/delivery/dto"
	"SSO/services"
	"SSO/user"
	"bytes"
	"errors"
	"html/template"
	"log"
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

func LoginHandler(auth services.AuthService) echo.HandlerFunc {
	return func(c *echo.Context) error {
		loginReq := dto.LoginRequest{}
		if err := c.Bind(&loginReq); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if err := loginReq.Validate(); err != nil {
			return c.JSON(http.StatusBadRequest, dto.LoginPage{
				Username: loginReq.Username,
				Error:    err.Error(),
			})
		}
		sessionIns, errLogin := auth.Login(loginReq.Username, loginReq.Password)
		if errLogin != nil {
			if errors.Is(errLogin, user.ErrInvalidCredentials) {
				return c.JSON(http.StatusUnauthorized, user.ErrInvalidCredentials)
			} else if errors.Is(errLogin, user.ErrNotFound) {
				return c.JSON(http.StatusUnauthorized, user.ErrInvalidCredentials)
			}
			log.Println(errLogin)
			return c.JSON(http.StatusInternalServerError, errors.New("internal server error"))
		}
		c.SetCookie(&http.Cookie{
			Name:     "SSO_session",
			Value:    sessionIns.ID,
			Path:     "/",
			MaxAge:   28800,
			Secure:   false,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})
		return c.Redirect(http.StatusSeeOther, "/me")
	}
}
