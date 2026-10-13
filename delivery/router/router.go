package router

import (
	"SSO/config"
	"SSO/delivery/handler"
	"log"

	"github.com/labstack/echo/v5"
)

func RouteResolver() {
	cfg := config.Load()

	e := echo.New()

	e.GET(".well-known/openid-configuration", handler.GetIssuerHander)

	log.Printf("server listening on port 8080 ... \n")

	log.Fatal(e.Start(":" + cfg.Port))
}
