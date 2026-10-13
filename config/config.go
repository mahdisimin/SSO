package config

import (
	"os"
	"strings"
)

type Config struct {
	Issuer string
	Port   string
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func Load() Config {
	issuer := getEnv("Issuer", "http://localhost:8080")
	port := getEnv("Port", "8080")
	strings.TrimSuffix(issuer, "/")
	config := Config{
		Issuer: issuer,
		Port:   port,
	}
	return config
}
