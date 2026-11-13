package main

import (
	"SSO/config"
	"SSO/delivery/router"
	"SSO/keys"
	"SSO/services"
	"SSO/user"
	"log"
)

func main() {
	err := config.LoadDotEnv(".env")
	if err != nil {
		log.Fatal("error on loading env variables : ", err)
	}
	cfg, confErr := config.Load()
	if confErr != nil {
		log.Fatal("error on loading config : ", confErr)
	}
	rsaKey, rsaLoadErr := keys.LoadOrGenerateKey(cfg.KeyPath)
	if rsaLoadErr != nil {
		log.Fatal("error on load / generate rsa key : ", rsaLoadErr)
	}
	Kid, errKeyID := keys.GenerateKID(&rsaKey.PublicKey)
	if errKeyID != nil {
		log.Fatal("error on generating kid : ", errKeyID)
	}
	issuer := services.NewIssuer(cfg.IssuerAddress)
	jwk := keys.NewJWK(rsaKey.PublicKey, Kid)
	jwks := keys.JWKs{Keys: []keys.JWK{*jwk}}

	newMemoryStore := user.NewMemoryStore()
	if cfg.TestUsername != "" && cfg.TestPassword != "" {
		NewUser := user.NewUser(cfg.TestUsername, "Test@gmail.com", cfg.TestPassword)
		if err := newMemoryStore.Create(NewUser); err != nil {
			log.Fatal("error on create user : ", err)
		}
	}

	router.RouteResolver(cfg.IssuerPort, issuer, jwks)

}
