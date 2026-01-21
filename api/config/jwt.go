package config

import "github.com/nixoncode/skillflow/pkg/env"

type JWTConfig struct {
	SecretKey      string
	ExpirationMins int
}

func loadJWTConfig() *JWTConfig {
	return &JWTConfig{
		SecretKey:      env.GetStringEnv("JWT_SECRET", ""),
		ExpirationMins: env.GetIntEnv("JWT_EXPIRATION_MINS", 15),
	}
}
