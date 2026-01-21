package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	App    *AppConfig
	DB     *DatabaseConfig
	Server *ServerConfig
}

func LoadConfig() *Config {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system/default environment variables")
	}

	return &Config{
		App:    loadAppConfig(),
		DB:     loadDatabaseConfig(),
		Server: loadServerConfig(),
	}
}
