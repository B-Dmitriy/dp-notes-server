package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"webservice/internal/db"

	"github.com/gorilla/mux"
	h "webservice/pgk/http"
)

const (
	defaultLimit = "10"
	defaultPage  = "1"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	page := r.FormValue("page")
	limit := r.FormValue("limit")

	if page == "" {
		page = defaultPage
	}

	if limit == "" {
		limit = defaultLimit
	}

	users, err := db.GetUsers(page, limit)
	if err != nil {
		log.Fatal(err)
	}

	err = h.WriteJSON(w, http.StatusOK, users)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)

	user, err := db.GetUserById(vars["user_id"])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user db.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}

	_ = db.CreateUser(user.Name, user.Email)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user db.User

	upVars := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}
	fmt.Printf("%+v \n", user)
	_ = db.UpdateUser(upVars["user_id"], user.Name, user.Email)
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)

	err := db.DeleteUserById(vars["user_id"])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(vars["user_id"])
}
