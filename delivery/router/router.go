package router

import (
	"SSO/delivery/handler"
	"SSO/keys"
	"SSO/services"
	"html/template"
	"log"

	"github.com/labstack/echo/v5"
)

func RouteResolver(port string, issuer services.IssuerStruct, jwks keys.JWKs, tmpl *template.Template) {

	e := echo.New()

	e.GET(".well-known/openid-configuration", handler.DefineIssuerHandler(issuer))
	e.GET("/.well-known/jwks.json", handler.JWKsHandler(jwks))
	e.GET("/login", handler.LoginPageHandler(tmpl))
	log.Printf("server listening on port %s ... \n", port)

	log.Fatal(e.Start(":" + port))
}
