package router

import (
	"SSO/delivery/handler"
	"SSO/keys"
	"SSO/services"
	"log"

	"github.com/labstack/echo/v5"
)

func RouteResolver(port string, issuer services.IssuerStruct, jwk keys.JWK) {

	e := echo.New()

	e.GET(".well-known/openid-configuration", handler.DefineIssuerHandler(issuer))
	e.GET("/.well-known/jwks.json", handler.JWKsHandler(jwk))

	log.Printf("server listening on port %s ... \n", port)

	log.Fatal(e.Start(":" + port))
}
