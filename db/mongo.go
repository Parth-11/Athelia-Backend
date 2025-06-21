package db

import (
	"log"
	"sync"

	"github.com/Parth-11/Athelia-Backend/configs"
	"gofr.dev/pkg/gofr/datasource/mongo"
)

var (
	once       sync.Once
	mongoStore *mongo.Client
)

func InitMongoStore() *mongo.Client {
	once.Do(func() {
		cfg := configs.NewConfig()

		mongoStore = mongo.New(mongo.Config{
			URI:               cfg.MongoURI,
			Database:          cfg.Database,
			ConnectionTimeout: cfg.Timeout,
		})

		log.Println("Initialized GoFr MongoDB store")
	})

	return mongoStore
}
