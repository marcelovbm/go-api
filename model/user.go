package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"e-mail" json:"e-mail"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}

func NewUser(newUser User) User {
	return User{primitive.NewObjectID(), newUser.Name, newUser.Email, time.Now()}
}
