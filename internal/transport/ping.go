package transport

import (
	"fmt"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: ping")
	fmt.Fprintf(w, "Ping!")
}
