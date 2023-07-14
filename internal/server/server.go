package server

import (
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Users\n")
}

func CommentsHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "comments\n")
}

func NewServer(host, port, endpoint string) error {
	r := mux.NewRouter().PathPrefix(endpoint).Subrouter()
	r.HandleFunc("/users", UsersHandler)
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
