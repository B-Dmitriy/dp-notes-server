package http

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) (err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err = json.NewEncoder(w).Encode(v)
	return
}
