// Package config provides application configuration loaded from environment variables.
// All settings have sensible defaults for local development.
package config

import "os"

// Config holds all application configuration.
type Config struct {
	Port       string
	JWTSecret  string
	DBPath     string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	RedisHost  string
	RedisPort  string
	KafkaBroker string
}

// Load returns a Config populated from environment variables,
// falling back to development-friendly defaults.
func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		JWTSecret:   getEnv("JWT_SECRET", "uaad-super-secret-key-2026"),
		DBPath:      getEnv("DB_PATH", "uaad.db"),
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "3306"),
		DBUser:      getEnv("DB_USER", "root"),
		DBPassword:  getEnv("DB_PASSWORD", ""),
		DBName:      getEnv("DB_NAME", "uaad"),
		RedisHost:   getEnv("REDIS_HOST", "localhost"),
		RedisPort:   getEnv("REDIS_PORT", "6379"),
		KafkaBroker: getEnv("KAFKA_BROKER", "localhost:9092"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
