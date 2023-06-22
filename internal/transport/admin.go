package transport

import (
	"html/template"
	"log"
	"net/http"
)

func AdminHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "Method not supported", http.StatusNotFound)
	}

	t, err := template.ParseFiles("../../templates/layout.html", "../../templates/admin.html")
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(writer, "props for admin")
	if err != nil {
		log.Fatal(err)
	}
}
