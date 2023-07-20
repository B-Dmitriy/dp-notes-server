package server

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"time"
)

func CommentsHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "comments\n")
}

func NewServer(host, port, endpoint string) error {
	r := mux.NewRouter().PathPrefix(endpoint).Subrouter()

	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{user_id}", GetUserById).Methods("GET")
	r.HandleFunc("/users/{user_id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{user_id}", DeleteUserById).Methods("DELETE")

	r.HandleFunc("/comments", CommentsHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    host + ":" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}
