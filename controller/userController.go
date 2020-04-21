package controller

import (
	"encoding/json"
	"net/http"

	"github.com/marcelovbm/go-api/db/dao"
	. "github.com/marcelovbm/go-api/model"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {}
func GetUser(w http.ResponseWriter, r *http.Request)  {}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user = NewUser(user)

	if err := dao.CreateUser(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, user)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {}
