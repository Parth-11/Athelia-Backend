package handlers

import (
    "go.mongodb.org/mongo-driver/bson"
    "gofr.dev/pkg/gofr"
    "myapp/models"
)

func Insert(ctx *gofr.Context) (any, error) {
    var p models.Person
    err := ctx.Bind(&p)
    if err != nil {
        return nil, err
    }

    res, err := ctx.Mongo.InsertOne(ctx, "collection", p)
    if err != nil {
        return nil, err
    }

    return res, nil
}

func Get(ctx *gofr.Context) (any, error) {
    var result models.Person

    p := ctx.Param("name")

    err := ctx.Mongo.FindOne(ctx, "collection", bson.D{{Key: "name", Value: p}}, &result)
    if err != nil {
        return nil, err
    }

    return result, nil
}