package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServiceName   string `env:"SERVICE_NAME"`
	ServerAddress string `env:"SERVER_ADDRESS"`
    HTTPPort      string `env:"HTTP_PORT"`
}

func Load() *Config {
    _ = godotenv.Load()

    return &Config{
        ServiceName:     getEnv("SERVICE_NAME"),
		ServerAddress:   getEnv("SERVER_ADDRESS"),
        HTTPPort:        getEnv("HTTP_PORT"),
    }
}

func getEnv(key string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return "NO_VALUE"
}