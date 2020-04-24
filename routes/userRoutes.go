package routes

import (
	"github.com/gorilla/mux"
	. "github.com/marcelovbm/go-api/controller"
)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/", GetUsers).Methods("GET")
	router.HandleFunc("/", CreateUser).Methods("POST")
	router.HandleFunc("/{id}", GetUser).Methods("GET")
	router.HandleFunc("/{id}", DeleteUser).Methods("DELETE")
}
