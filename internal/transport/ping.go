package transport

import (
	"fmt"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: ping")
	_, err := fmt.Fprintf(w, "Ping!")
	if err != nil {
		return
	}
}
