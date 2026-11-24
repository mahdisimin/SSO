package router

import (
	"SSO/config"
	"SSO/delivery/handler"
	authHandler "SSO/delivery/handler/auth"
	mv "SSO/delivery/middlerware"
	"SSO/keys"
	"SSO/services"
	authSvc "SSO/services/auth"
	"html/template"
	"log"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func RouteResolver(cfg config.Config, issuer services.IssuerStruct, jwks keys.JWKs, tmpl *template.Template, auth authSvc.AuthService) {

	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(mv.ServedBy(cfg.InstanceID))
	e.GET(".well-known/openid-configuration", handler.DefineIssuerHandler(issuer))
	e.GET("/.well-known/jwks.json", handler.JWKsHandler(jwks))
	e.GET("/login", authHandler.LoginPageHandler(tmpl))
	e.POST("/login", authHandler.LoginHandler(auth, cfg, tmpl))
	e.GET("/me", authHandler.CheckSessionHandler(auth, tmpl))
	e.POST("/logout", authHandler.LogoutHandler(auth, cfg))
	log.Printf("server listening on port %s ... \n", cfg.IssuerPort)

	log.Fatal(e.Start(":" + cfg.IssuerPort))
}
