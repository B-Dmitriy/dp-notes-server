package transport

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Statistic map[string]string

func MonitoringHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "Method not supported", http.StatusNotFound)
	}

	t, err := template.ParseFiles("../../templates/layout.html", "../../templates/monitoring.html")
	if err != nil {
		log.Fatal(err)
	}

	tagsByteSlice, _ := os.ReadFile("../../data/data.json")

	var tagsJson Tags

	err = json.Unmarshal(tagsByteSlice, &tagsJson)

	if err != nil {
		log.Fatal(err)
	}

	stat := Statistic{
		"TagsCount": "",
	}

	stat["TagsCount"] = strconv.Itoa(len(tagsJson.Tags))

	err = t.Execute(writer, stat)
	if err != nil {
		log.Fatal(err)
	}
}
