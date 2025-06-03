package models

type user struct {
	Name string `bson:"name" json:"name"`
	Age  int    `bson:"age" json:"age"`
}
