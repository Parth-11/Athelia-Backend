package main

import (
	"github.com/Parth-11/Athelia-Backend/db"
	"github.com/Parth-11/Athelia-Backend/routes"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	app.GET("/greet", func(context *gofr.Context) (any, error) { return "Hello World!", nil })

	mongoClient := db.InitMongoStore()

	app.AddMongo(mongoClient)

	routes.Setup_user_routes(app)

	app.Run()
}
