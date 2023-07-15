package server

import (
	"fmt"
	"net/http"
	"webservice/internal/db"

	"github.com/gorilla/mux"
)

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)

	user, _ := db.GetUserById(vars["user_id"])

	fmt.Println(user)
}
