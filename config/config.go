package config

import (
	"os"
	"strings"
)

type Config struct {
	IssuerAddress string
	IssuerPort    string
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func Load() Config {
	SSO_ISSUER := getEnv("Issuer", "http://localhost")
	SSO_PORT := getEnv("Port", "8080")
	SSO_ISSUER = strings.TrimSuffix(SSO_ISSUER, "/")
	config := Config{
		IssuerAddress: SSO_ISSUER,
		IssuerPort:    SSO_PORT,
	}
	return config
}
