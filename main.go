package main

import (
	"SSO/config"
	"SSO/delivery/router"
	"SSO/keys"
	"SSO/services"
	"SSO/services/auth"
	"SSO/session"
	"SSO/user"
	"html/template"
	"log"

	"github.com/redis/go-redis/v9"
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

	DBStore := user.NewMemoryStore()

	memoryStore := newMemoryStore(cfg)

	if cfg.TestUsername != "" && cfg.TestPassword != "" {
		NewUser := user.NewUser(cfg.TestUsername, "Test@gmail.com", cfg.TestPassword)
		if cfg.TestUserID != "" {
			NewUser.ID = cfg.TestUserID
		}
		if err := DBStore.Create(NewUser); err != nil {
			log.Fatal("error on create user : ", err)
		}
	}
	tmpl, err := template.ParseGlob("web/template/*.html")
	if err != nil {
		log.Fatal("error on parsing templates : ", err)
	}
	auth := auth.NewAuthService(DBStore, memoryStore, cfg.SessionTTL)
	router.RouteResolver(cfg, issuer, jwks, tmpl, auth)

}

func newMemoryStore(cfg config.Config) session.Store {
	switch cfg.SessionStore {
	case "redis":
		{
			client := redis.NewClient(&redis.Options{
				Addr:     cfg.RedisHost,
				Password: "",
				DB:       0,
			},
			)
			rs := session.NewRedis(client)
			if err := rs.Client.Ping(rs.Ctx).Err(); err != nil {
				log.Fatal("error on connecting to redis : ", err)
			}
			return rs
		}
	case "memory":
		return session.NewMemoryStore()
	default:
		log.Fatal("invalid session store")
	}
	return nil
}
