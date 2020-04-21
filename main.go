package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	. "github.com/marcelovbm/go-api/db"
)

var db = Connection{}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db.URI = os.Getenv("URI")
	db.Connect()
}

func main() {
	log.Println("Application is UP!")
}
