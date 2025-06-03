package main

import (
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/datasource/mongo"
)

func main() {
	app := gofr.New()

	app.GET("/greet", func(context *gofr.Context) (any, error) { return "Hello World!", nil })

	db := mongo.New(mongo.Config{})

	app.AddMongo(db)

	app.Run()
}
