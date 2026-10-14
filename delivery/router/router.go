package router

import (
	"SSO/delivery/handler"
	"SSO/services"
	"log"

	"github.com/labstack/echo/v5"
)

func RouteResolver(port string, issuer services.IssuerStruct) {

	e := echo.New()

	e.GET(".well-known/openid-configuration", handler.DefineIssuerHandler(issuer))

	log.Printf("server listening on port %s ... \n", port)

	log.Fatal(e.Start(":" + port))
}
