package models

type Person struct {
    Name string `bson:"name" json:"name"`
    Age  int    `bson:"age" json:"age"`
    City string `bson:"city" json:"city"`
}