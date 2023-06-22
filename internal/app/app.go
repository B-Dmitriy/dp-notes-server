package app

import (
	"dpnotes/internal/transport"
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/admin", transport.AdminHandler)
	http.HandleFunc("/monitoring", transport.MonitoringHandler)
	http.HandleFunc("/api/ping", transport.PingHandler)
	http.HandleFunc("/api/tags/", transport.TagsHandler)

	err := http.ListenAndServe("localhost:8011", nil)
	log.Fatal(err)
}
