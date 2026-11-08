package config

import (
	"os"
	"strings"
)

type Config struct {
	IssuerAddress string
	IssuerPort    string
	KeyPath       string
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func Load() Config {
	SSO_ISSUER := getEnv("SSO_ISSUER", "http://localhost:8080")
	SSO_PORT := getEnv("SSO_PORT", "8080")
	SSO_KEY_PATH := getEnv("SSO_KEY_PATH", "./secrets/signing.pem")
	SSO_ISSUER = strings.TrimSuffix(SSO_ISSUER, "/")
	config := Config{
		IssuerAddress: SSO_ISSUER,
		IssuerPort:    SSO_PORT,
		KeyPath:       SSO_KEY_PATH,
	}
	return config
}
