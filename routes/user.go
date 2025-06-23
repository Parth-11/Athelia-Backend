package routes

import (
	"github.com/Parth-11/Athelia-Backend/handler"
	"gofr.dev/pkg/gofr"
)

func Setup_user_routes(app *gofr.App) {
	app.GET("/user{id}", handler.Get_user)
}
