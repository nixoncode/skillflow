package config

import "github.com/nixoncode/skillflow/pkg/env"

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func loadDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     env.GetStringEnv("DB_HOST", "localhost"),
		Port:     env.GetIntEnv("DB_PORT", 3306),
		User:     env.GetStringEnv("DB_USER", ""),
		Password: env.GetStringEnv("DB_PASSWORD", ""),
		Name:     env.GetStringEnv("DB_NAME", "skillflow_db"),
	}
}
