package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	. "github.com/marcelovbm/go-api/controller"
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

	router := mux.NewRouter()

	router.HandleFunc("/user", GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", GetUser).Methods("GET")
	router.HandleFunc("/user", CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
	log.Println("Application is UP!")
}
