package auth

import (
	"SSO/delivery/dto"
	"SSO/services/auth"
	"bytes"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v5"
)

func CheckSessionHandler(auth auth.AuthService, tmpl *template.Template) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var buf bytes.Buffer
		cookieIns, errCookie := c.Cookie(cookieName)
		if errCookie != nil {
			return c.Redirect(http.StatusFound, "/login") //Check if cookie does not exist , then user will be redirected to login page
		}
		sessionID := cookieIns.Value
		ok := auth.CheckSession(sessionID)
		if !ok {
			return c.Redirect(http.StatusFound, "/login")
		}

		if err := tmpl.ExecuteTemplate(&buf, "me.html", dto.MePage{Username: "MMAD"}); err != nil {
			return err
		}
		return c.HTML(http.StatusOK, buf.String())
	}

}
