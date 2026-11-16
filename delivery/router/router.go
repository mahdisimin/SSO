package router

import (
	"SSO/config"
	"SSO/delivery/handler"
	"SSO/keys"
	"SSO/services"
	"SSO/services/auth"
	"html/template"
	"log"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func RouteResolver(cfg config.Config, issuer services.IssuerStruct, jwks keys.JWKs, tmpl *template.Template, auth auth.AuthService) {

	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.GET(".well-known/openid-configuration", handler.DefineIssuerHandler(issuer))
	e.GET("/.well-known/jwks.json", handler.JWKsHandler(jwks))
	e.GET("/login", handler.LoginPageHandler(tmpl))
	e.POST("/login", handler.LoginHandler(auth, cfg, tmpl))
	e.GET("/me", handler.CheckSessionHandler(auth))
	log.Printf("server listening on port %s ... \n", cfg.IssuerPort)

	log.Fatal(e.Start(":" + cfg.IssuerPort))
}
