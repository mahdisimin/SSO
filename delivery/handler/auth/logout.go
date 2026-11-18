package auth

import (
	"SSO/config"
	"SSO/services/auth"
	"net/http"

	"github.com/labstack/echo/v5"
)

func LogoutHandler(authSvc auth.AuthService, cfg config.Config) echo.HandlerFunc {
	return func(c *echo.Context) error {
		cookieIns, errCookie := c.Cookie(cookieName)
		if errCookie != nil {
			return c.Redirect(http.StatusSeeOther, "/login")
		}
		sessionID := cookieIns.Value
		if errLogout := authSvc.Logout(sessionID); errLogout != nil {
			return errLogout
		}
		c.SetCookie(&http.Cookie{
			Name:     cookieName,
			Value:    sessionID,
			Path:     "/",
			MaxAge:   -1,
			Secure:   cfg.CookieSecure,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})
		return c.Redirect(http.StatusSeeOther, "/login")
	}
}
