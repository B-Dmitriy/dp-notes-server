package app

import (
	"dpnotes/internal/transport"
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/ping", transport.PingHandler)
	http.HandleFunc("/tags/", transport.TagsHandler)

	err := http.ListenAndServe("localhost:8011", nil)
	log.Fatal(err)
}
