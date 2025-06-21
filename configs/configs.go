package configs

import (
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
		MongoURI:    "mongodb://localhost:27017",
		Database:    "test",
		Timeout:     4 * time.Second,
		MaxPoolSize: 20,
	}
}
