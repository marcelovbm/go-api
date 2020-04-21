package dao

import (
	"context"

	. "github.com/marcelovbm/go-api/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// DATABASE INSTANCE
var collection *mongo.Collection

func UsersCollection(c *mongo.Database) {
	collection = c.Collection("users")
}

const (
	COLLECTION = "users"
)

func CreateUser(user User) error {
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}
