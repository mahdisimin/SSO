package main

import (
	"SSO/config"
	"SSO/delivery/router"
	"SSO/jwt"
	"SSO/keys"
	"SSO/services"
	"fmt"
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
	fmt.Println(rsaKey.N.BitLen())
	Kid, errKeyID := keys.GenerateKID(&rsaKey.PublicKey)
	if errKeyID != nil {
		log.Fatal("error on generating kid : ", errKeyID)
	}
	fmt.Println(Kid)
	signer := jwt.NewSigner(rsaKey, Kid)
	_ = signer
	issuer := services.NewIssuer(cfg.IssuerAddress)
	jwk := keys.NewJWK(rsaKey.PublicKey, Kid)
	jwks := keys.JWKs{Keys: []keys.JWK{*jwk}}
	router.RouteResolver(cfg.IssuerPort, issuer, jwks)

}
