package server

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Ping\n")
}

func NewServer(host, port, endpoint string) error {
	addr := host + ":" + port
	fmt.Println(host, port, endpoint)
	http.HandleFunc("/", handler)
	return http.ListenAndServe(addr, nil)
}
