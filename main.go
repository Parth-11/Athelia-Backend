package main

import (
	"github.com/Parth-11/Athelia-Backend/db"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	app.GET("/greet", func(context *gofr.Context) (any, error) { return "Hello World!", nil })

	mongoClient := db.InitMongoStore()

	app.AddMongo(mongoClient)

	app.Run()
}
