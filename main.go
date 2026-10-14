package main

import (
	"SSO/config"
	"SSO/delivery/router"
	"SSO/services"
	"log"
)

func main() {
	err := config.LoadDotEnv(".env")
	if err != nil {
		log.Fatal("error on loading env variables : ", err)
	}
	cfg := config.Load()
	issuer := services.NewIssuer(cfg.IssuerAddress)

	router.RouteResolver(cfg.IssuerPort, issuer)
}
