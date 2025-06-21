package handler

import (
	"github.com/Parth-11/Athelia-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"gofr.dev/pkg/gofr"
)

func Get_user(ctx *gofr.Context) (any, error) {
	var res models.User

	name := ctx.Param("name")
	err := ctx.Mongo.FindOne(ctx, "collection", bson.D{{Key: "name", Value: name}}, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
