package server

import (
	"fmt"
	"log"
	"net/http"
	"webservice/internal/db"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	page := r.FormValue("page")

	users, err := db.GetUsers(page)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)
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

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)

	err := db.DeleteUserById(vars["user_id"])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(vars["user_id"])
}
