package config

import "github.com/nixoncode/skillflow/pkg/env"

type AppConfig struct {
	IsDebug  bool
	Name     string
	Version  string
	LogLevel string
}

func loadAppConfig() *AppConfig {
	return &AppConfig{
		IsDebug:  env.GetBoolEnv("APP_DEBUG", true),
		Name:     env.GetStringEnv("APP_NAME", "Skillflow"),
		Version:  env.GetStringEnv("APP_VERSION", "1.0.0"),
		LogLevel: env.GetStringEnv("APP_LOG_LEVEL", "debug"),
	}
}
