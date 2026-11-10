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
	cfg := config.Load()
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

	newStore := user.TestStore{}
	newUser := user.NewUser(cfg.TestUsername, "Test@gmail.com", cfg.TestPassword)
	if _, err := newStore.Create(newUser); err != nil {
		panic(err)
	}

	router.RouteResolver(cfg.IssuerPort, issuer, jwks)

}
