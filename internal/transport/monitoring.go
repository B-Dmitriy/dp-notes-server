package transport

import (
	"html/template"
	"log"
	"net/http"
)

func MonitoringHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "Method not supported", http.StatusNotFound)
	}

	t, err := template.ParseFiles("../../templates/layout.html", "../../templates/monitoring.html")
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(writer, "Title monitoring")
	if err != nil {
		log.Fatal(err)
	}
}
