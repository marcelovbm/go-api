package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	. "github.com/marcelovbm/go-api/db"
	. "github.com/marcelovbm/go-api/routes"
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

	adress := ":" + os.Getenv("PORT")
	router := mux.NewRouter()

	userSubRouter := router.PathPrefix("/users").Subrouter()
	UserRoutes(userSubRouter)

	log.Fatal(http.ListenAndServe(adress, router))
	log.Println("Application is UP!")
}
