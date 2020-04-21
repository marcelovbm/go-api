package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/marcelovbm/go-api/db/dao"
)

type Connection struct {
	URI string
}

func (c *Connection) Connect() {

	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(c.URI))

	if err != nil {
		log.Fatal(err)
	}
	// Disconect
	defer cancel()

	// Check the connection
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	db := client.Database("go-api")
	dao.UsersCollection(db)

	return
}
