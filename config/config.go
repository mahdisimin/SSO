package config

import (
	"errors"
	"os"
	"strings"
	"time"
)

type Config struct {
	IssuerAddress string
	IssuerPort    string
	KeyPath       string
	TestUsername  string
	TestPassword  string
	SessionTTL    time.Duration
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func Load() (Config, error) {
	SSO_ISSUER := getEnv("SSO_ISSUER", "http://localhost:8081")
	SSO_PORT := getEnv("SSO_PORT", "8081")
	SSO_KEY_PATH := getEnv("SSO_KEY_PATH", "./secrets/signing.pem")
	SSO_SEED_USERNAME := getEnv("SSO_SEED_USERNAME", "")
	SSO_SEED_PASSWORD := getEnv("SSO_SEED_PASSWORD", "")
	SSO_SESSION_TTL := getEnv("SSO_SESSION_TTL", "8h")
	SSO_ISSUER = strings.TrimSuffix(SSO_ISSUER, "/")

	SessionTTL, err := time.ParseDuration(SSO_SESSION_TTL)
	if err != nil {
		return Config{}, errors.New("invalid SSO_SESSION_TTL")

	}
	config := Config{
		IssuerAddress: SSO_ISSUER,
		IssuerPort:    SSO_PORT,
		KeyPath:       SSO_KEY_PATH,
		TestUsername:  SSO_SEED_USERNAME,
		TestPassword:  SSO_SEED_PASSWORD,
		SessionTTL:    SessionTTL,
	}
	return config, nil
}
