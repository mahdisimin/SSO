package handler

import (
	"SSO/config"
	"SSO/services"
	"net/http"

	"github.com/labstack/echo/v5"
)

func GetIssuerHander(c *echo.Context) error {
	cfg := config.Load()
	issuer := services.NewIssuer(cfg.Issuer)
	return c.JSON(http.StatusOK, issuer)
}
