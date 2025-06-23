package configs

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	MongoURI    string
	Database    string
	Timeout     time.Duration
	MaxPoolSize uint64
}

func NewConfig() *Config {
	return &Config{
		MongoURI:    getEnv("MONGO_URI", "mongodb://localhost:27017"),
		Database:    getEnv("MONGO_DB", "test"),
		Timeout:     getEnvAsDuration("MONGO_TIMEOUT", 4*time.Second),
		MaxPoolSize: getEnvAsUint64("MONGO_POOL_SIZE", 20),
	}
}

func getEnv(key string, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultVal
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		d, err := time.ParseDuration(value)
		if err == nil {
			return d
		}
	}

	return defaultValue
}

func getEnvAsUint64(key string, defaultValue uint64) uint64 {
	if value := os.Getenv(key); value != "" {
		parsed, err := strconv.ParseUint(value, 10, 64)

		if err == nil {
			return parsed
		}
	}

	return defaultValue
}
