package handler

import (
	"SSO/config"
	"SSO/delivery/dto"
	"SSO/services"
	"SSO/user"
	"bytes"
	"errors"
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

func LoginHandler(auth services.AuthService, conf config.Config, tmpl *template.Template) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var buf bytes.Buffer
		loginReq := dto.LoginRequest{}
		if err := c.Bind(&loginReq); err != nil {
			return err
		}
		if err := loginReq.Validate(); err != nil {
			if errTmp := tmpl.ExecuteTemplate(&buf, "login.html", dto.LoginPage{
				Username: loginReq.Username,
				Error:    user.ErrInvalidCredentials.Error(),
			}); errTmp != nil {
				return errTmp
			}
			return c.HTML(http.StatusBadRequest, buf.String())
		}
		sessionIns, errLogin := auth.Login(loginReq.Username, loginReq.Password)
		if errLogin != nil {
			if errors.Is(errLogin, user.ErrInvalidCredentials) || errors.Is(errLogin, user.ErrNotFound) {
				if errTmp := tmpl.ExecuteTemplate(&buf, "login.html", dto.LoginPage{
					Username: loginReq.Username,
					Error:    user.ErrInvalidCredentials.Error(),
				}); errTmp != nil {
					return errTmp
				}
				return c.HTML(http.StatusUnauthorized, buf.String())
			}
			return errLogin
		}
		c.SetCookie(&http.Cookie{
			Name:     "SSO_session",
			Value:    sessionIns.ID,
			Path:     "/",
			MaxAge:   int(time.Until(sessionIns.ExpiresAt).Seconds()),
			Secure:   conf.CookieSecure,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})
		return c.Redirect(http.StatusSeeOther, "/me")
	}
}
