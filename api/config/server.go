package config

import "github.com/nixoncode/skillflow/pkg/env"

type ServerConfig struct {
	Host               string
	Port               int
	CORSAllowedOrigins string
}

func loadServerConfig() *ServerConfig {
	return &ServerConfig{
		Host:               env.GetStringEnv("SERVER_HOST", "localhost"),
		Port:               env.GetIntEnv("SERVER_PORT", 8080),
		CORSAllowedOrigins: env.GetStringEnv("SERVER_CORS_ALLOWED_ORIGINS", "*"),
	}
}
