package dao

import (
	"context"
	"log"

	. "github.com/marcelovbm/go-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetAllUsers() []User {
	var users []User
	cursor, err := collection.Find(context.TODO(), bson.D{}, options.Find())
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {

		var user User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cursor.Close(context.TODO())

	return users
}
