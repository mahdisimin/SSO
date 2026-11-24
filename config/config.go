package config

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	IssuerAddress string
	IssuerPort    string
	KeyPath       string
	TestUsername  string
	TestPassword  string
	TestUserID    string
	SessionTTL    time.Duration
	CookieSecure  bool
	SessionStore  string
	RedisHost     string
	InstanceID    string
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
	SSO_COOKIE_SECURE := getEnv("SSO_COOKIE_SECURE", "false")
	SSO_ISSUER = strings.TrimSuffix(SSO_ISSUER, "/")
	SSO_SESSION_STORE := getEnv("SSO_SESSION_STORE", "memory")
	SSO_REDIS_ADDR := getEnv("SSO_REDIS_ADDR", "localhost:6379")
	SSO_SEED_USERID := getEnv("SSO_SEED_USERID", "")
	SSO_INSTANCE_ID := getEnv("SSO_INSTANCE_ID", "sso")

	SessionTTL, err := time.ParseDuration(SSO_SESSION_TTL)
	if err != nil {
		return Config{}, errors.New("invalid SSO_SESSION_TTL")

	}
	ssoSecureBool, errorConv := strconv.ParseBool(SSO_COOKIE_SECURE)
	if errorConv != nil {
		return Config{}, errors.New("invalid SSO_COOKIE_SECURE")
	}

	config := Config{
		IssuerAddress: SSO_ISSUER,
		IssuerPort:    SSO_PORT,
		KeyPath:       SSO_KEY_PATH,
		TestUsername:  SSO_SEED_USERNAME,
		TestPassword:  SSO_SEED_PASSWORD,
		TestUserID:    SSO_SEED_USERID,
		SessionTTL:    SessionTTL,
		CookieSecure:  ssoSecureBool,
		SessionStore:  SSO_SESSION_STORE,
		RedisHost:     SSO_REDIS_ADDR,
		InstanceID:    SSO_INSTANCE_ID,
	}
	return config, nil
}
