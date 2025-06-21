package configs

import (
	"os"
	"time"
)

type Config struct {
	MongoURI    string
	Database    string
	MaxPoolSize uint64
	Timeout     time.Duration
}

func NewConfig() *Config {
	return &Config{
		MongoURI:    getEnv("MONGO_URI", ""),
		Database:    getEnv("DATABASE_NAME", ""),
		MaxPoolSize: 100,
		Timeout:     10 * time.Second,
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
