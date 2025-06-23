package models

import "github.com/gofrs/uuid/v5"

type User struct {
	UserId    uuid.UUID `bson:"user_id" json:"id"`
	FirstName string    `bson:"first_name" json:"first_name"`
	LastName  string    `bson:"last_name" json:"last_name"`
	Email     string    `bson:"email" json:"email"`
	Phone     string    `bson:"phone" json:"phone"`
}
