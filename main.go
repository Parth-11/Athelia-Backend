package main

import (
	"github.com/Parth-11/Athelia-Backend/db"
	"gofr.dev/pkg/gofr"
)

func main() {
    app := gofr.New()

    config.SetupMongo(app)

	mongoClient := db.InitMongoStore()

	app.AddMongo(mongoClient)

    app.Run()
}