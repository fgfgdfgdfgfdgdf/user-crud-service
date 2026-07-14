package config

import (
	"os"
	"strconv"
)

type APIConfig struct {
	API_LISTEN_ADDR string
	API_PORT        int
}

type PostgresConfig struct {
	POSTGRES_USER     string
	POSTGRES_PORT     int
	POSTGRES_HOST     string
	POSTGRES_DB       string
	POSTGRES_PASSWORD string
}

var (
	POSTGRES PostgresConfig
	API      APIConfig
)

func Init() error {
	API.API_LISTEN_ADDR = getEnv("API_LISTEN_ADDR", "0.0.0.0")
	API.API_PORT = getEnvInt("API_PORT", 8080)

	POSTGRES.POSTGRES_USER = getEnv("POSTGRES_USER", "")
	POSTGRES.POSTGRES_PASSWORD = getEnv("POSTGRES_PASSWORD", "")
	POSTGRES.POSTGRES_HOST = getEnv("POSTGRES_HOST", "")
	POSTGRES.POSTGRES_DB = getEnv("POSTGRES_DB", "")
	POSTGRES.POSTGRES_PORT = getEnvInt("POSTGRES_PORT", 5432)

	return nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	valStr := getEnv(key, "")
	if valStr == "" {
		return defaultValue
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultValue
	}
	return val
}
