package main

import (
    "gofr.dev/pkg/gofr"
    "myapp/config"
    "myapp/handlers"
)

func main() {
    app := gofr.New()

    config.SetupMongo(app)

    app.POST("/users", handlers.Insert)
    app.GET("/users", handlers.Get)

    app.Run()
}