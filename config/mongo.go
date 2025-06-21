package config

import (
    "time"
    "gofr.dev/pkg/gofr"
    "gofr.dev/pkg/gofr/datasource/mongo"
)

func SetupMongo(app *gofr.App) {
    db := mongo.New(mongo.Config{
        URI:               "mongodb://localhost:27017",
        Database:          "test",
        ConnectionTimeout: 4 * time.Second,
    })
    
    app.AddMongo(db)
}